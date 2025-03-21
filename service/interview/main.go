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
}

// Initialize gRPC server
func initializeGrpc() {
	grpcPort := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:"+grpcPort))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	interviewServiceImpl := &api.InterviewServiceServerImpl{
		Database: <-databaseChannel,
	}
	api.RegisterInterviewServiceServer(grpcServer, interviewServiceImpl)
	grpcServer.Serve(lis)
}
