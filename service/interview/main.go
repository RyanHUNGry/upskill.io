// Interview service driver

package main

import (
	"context"
	"flag"
	"fmt"
	"interview/src/api"
	"interview/src/db"
	"interview/src/llm"
	"interview/src/producer"
	"interview/src/utils"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

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
	modelChannel    chan *llm.Model                 = make(chan *llm.Model)
	sigIntChannel   chan os.Signal                  = make(chan os.Signal, 1)
)

func main() {
	// Trap SIGINT to gracefully shutdown, propagating signal via context to other goroutines
	rootCtx := context.Background()
	signal.Notify(sigIntChannel, syscall.SIGINT)
	ctx, cancel := context.WithCancel(rootCtx)

	// Load environment variables
	envFileName := "interview.env"
	envFilePath := filepath.Join(utils.GetWorkingDirectory(), envFileName)
	err := godotenv.Load(envFilePath)

	if err != nil {
		log.Fatalf("Error loading %s file", envFileName)
	}

	go initializeKafka(ctx)
	go initializeGrpc(ctx)
	go initializeCassandra(ctx)
	go initializeModel(ctx)

	// Under the hood, cancel() calls close(ctx.Done()) which causes <-ctx.Done() to return a value immediately
	<-sigIntChannel
	cancel()
	fmt.Println("Shutting down gracefully...")
	time.Sleep(750 * time.Millisecond)
}

// Initialize OpenAI model
func initializeModel(ctx context.Context) {
	model := llm.InitializeModel(ctx)
	modelChannel <- model
	fmt.Println("Model initialized ✅")
}

// Initialize Kafka producer
func initializeKafka(ctx context.Context) {
	// Kafka in Docker runs with latest 4.0.0 image version
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

	select {
	case <-ctx.Done():
		return
	}
}

// Initialize Cassandra connection
func initializeCassandra(ctx context.Context) {
	cassandraHost := os.Getenv("CASSANDRA_HOST")
	cassandraPort := os.Getenv("CASSANDRA_PORT")

	db, err := db.Connect(cassandraHost, cassandraPort, ctx)
	if err != nil {
		log.Fatal("Database connection failed")
	}
	defer db.Session.Close()
	databaseChannel <- db

	clearDb := flag.Bool("c", false, "Clear all tables")
	flag.Parse()
	if *clearDb {
		db.DropAllTables()
	}

	err = db.InitializeTables()
	if err != nil {
		log.Fatalf("Error initializing tables: %v", err)
	}
	fmt.Println("Cassandra initialized ✅")

	select {
	case <-ctx.Done():
		return
	}
}

// Initialize gRPC server
func initializeGrpc(ctx context.Context) {
	grpcPort := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:"+grpcPort))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	logger := kitlog.NewLogfmtLogger(os.Stdout)
	asyncLogGenerator := <-producerChannel

	// From the logging thread, producer will send success or failiure messages to these logging threads
	go func(ctx context.Context) {
		for {
			select {
			case producerSuccessMessage := <-asyncLogGenerator.Producer.Successes():
				log.Printf("Produced message to topic %s at offset %d\n", producerSuccessMessage.Topic, producerSuccessMessage.Offset)
			case <-ctx.Done():
				return
			}
		}
	}(ctx)
	go func(ctx context.Context) {
		for {
			select {
			case producerErrorMessage := <-asyncLogGenerator.Producer.Errors():
				log.Printf("Failed to produce message: %v\n", producerErrorMessage.Err)
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

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
		})
	}

	// For verbosity reasons as defaults options log gRPC from start to finish
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
		Model:    <-modelChannel,
	}
	api.RegisterInterviewServiceServer(grpcServer, interviewServiceImpl)
	fmt.Println("gRPC Server initialized ✅")

	go func() {
		defer grpcServer.GracefulStop()
		<-ctx.Done()
		return
	}()

	grpcServer.Serve(lis)
}
