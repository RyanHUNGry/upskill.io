// Interview service driver

package main

import (
	"context"
	"flag"
	"fmt"
	"interview/src/api"
	"interview/src/db"
	"interview/src/db/table"
	"interview/src/producer"
	"interview/src/utils"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/IBM/sarama"
	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	databaseChannel chan *db.Database               = make(chan *db.Database)
	producerChannel chan producer.AsyncLogGenerator = make(chan producer.AsyncLogGenerator)
	blockingChannel chan *struct{}                  = make(chan *struct{})
)

func main() {
	envFileName := "interview.env"
	envFilePath := filepath.Join(utils.GetWorkingDirectory(), envFileName)
	err := godotenv.Load(envFilePath)

	if err != nil {
		log.Fatalf("Error loading %s file", envFileName)
	}

	go initializeKafka()
	go initializeGrpc()
	go initializeCassandra()

	<-blockingChannel
}

func initializeKafka() {
	// Kafka in Docker runs with latest 4.0.0 version
	kafkaVersion := sarama.MaxVersion
	kafkaBrokers := os.Getenv("KAFKA_PEERS")
	brokerList := strings.Split(kafkaBrokers, ",")

	asyncLogGenerator, err := producer.InitializeAsyncLogGenerator(brokerList, kafkaVersion)
	producerChannel <- asyncLogGenerator

	if err != nil {
		log.Fatalf("Error initializing Kafka producer: %v", err)
	}

	defer asyncLogGenerator.Close()

	fmt.Println("Kafka initialized ✅")
	<-blockingChannel
}

func initializeCassandra() {
	cassandraHost := os.Getenv("CASSANDRA_HOST")
	cassandraPort := os.Getenv("CASSANDRA_PORT")

	dbContext := context.Background()
	db, err := db.Connect(cassandraHost, cassandraPort, dbContext)
	if err != nil {
		log.Fatal("Database connection failed")
	}
	defer db.Session.Close()
	databaseChannel <- db

	clearDb := flag.Bool("c", false, "Clear all tables")
	flag.Parse()
	if *clearDb {
		table.DropAllTables(db.Session, db.Ctx)
		return
	}

	// Initialize table only for an empty database
	table.InitializeTables(db.Session, db.Ctx)
	fmt.Println("Cassandra initialized ✅")
	<-blockingChannel
}

// Initialize gRPC server
func initializeGrpc() {
	grpcPort := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:"+grpcPort))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	asyncLogGenerator := <-producerChannel
	logger := kitlog.NewLogfmtLogger(os.Stdout)
	interceptorLogger := func(l kitlog.Logger) logging.Logger {
		// The logging.LoggerFunc type is a typed function implementing logging.Logger interface so type convert ordinary function to be of logging.LoggerFunc
		return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
			largs := append([]any{"msg", msg}, fields...)
			switch lvl {
			case logging.LevelDebug:
				_ = level.Debug(l).Log(largs...)
			case logging.LevelInfo:
				_ = level.Info(l).Log(largs...)
			case logging.LevelWarn:
				_ = level.Warn(l).Log(largs...)
			case logging.LevelError:
				_ = level.Error(l).Log(largs...)
			default:
				panic(fmt.Sprintf("unknown level %v", lvl))
			}

			var grpcStartTime string
			for i, field := range fields {
				if field.(string) == "grpc.start_time" {
					grpcStartTime = fields[i+1].(string)
					break
				}
			}
			var keyEncoder sarama.StringEncoder = sarama.StringEncoder(grpcStartTime)
			var valueEncoder sarama.StringEncoder = sarama.StringEncoder(fmt.Sprintf("%v", largs))

			asyncLogGenerator.Producer.Input() <- &sarama.ProducerMessage{
				Topic: "interview_service_logs",
				Key:   keyEncoder,
				Value: valueEncoder,
			}

			// figure out how to get the producer errors that are generated in another goroutine
			// we should log the error here, but keep the thread alive
			// producerErr := <-asyncLogGenerator.Producer.Errors()
			// fmt.Println(producerErr.Err)
		})
	}

	// For verbosity reasons only as defaults options log gRPC from start to finish
	var loggingOption []logging.Option = []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}
	var serverOption []grpc.ServerOption = []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(interceptorLogger(logger), loggingOption...),
		),
	}

	grpcServer := grpc.NewServer(serverOption...)
	interviewServiceImpl := &api.InterviewServiceServerImpl{
		Database: <-databaseChannel,
	}
	api.RegisterInterviewServiceServer(grpcServer, interviewServiceImpl)
	fmt.Println("gRPC Server initialized ✅")
	grpcServer.Serve(lis)
}
