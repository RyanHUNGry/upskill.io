// Interview service driver

package main

import (
	"context"
	"flag"
	"fmt"
	"interview/src/api"
	"interview/src/db"
	"interview/src/db/table"
	"interview/src/utils"
	"log"
	"net"
	"os"
	"path/filepath"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	databaseChannel chan *db.Database = make(chan *db.Database)
	blockingChannel chan *struct{}    = make(chan *struct{})
)

func main() {
	envFileName := "interview.env"
	envFilePath := filepath.Join(utils.GetWorkingDirectory(), envFileName)
	err := godotenv.Load(envFilePath)

	if err != nil {
		log.Fatalf("Error loading %s file", envFileName)
	}

	go initializeGrpc()
	go initializeCassandra()

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
		})
	}

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
