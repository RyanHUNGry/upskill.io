package main

import (
	"context"
	"fmt"
	"net"
	"os"

	pb "interview-service/api"
	"interview-service/internal/db"
	"interview-service/internal/server"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load() // Default .env file is loaded if no file is specified

	if err != nil {
		panic(err)
	}

	port := os.Getenv("INTERVIEW_SERVICE_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))

	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	var answerServiceServer *server.AnswerServiceServer = new(server.AnswerServiceServer) // same as &server.AnswerServiceServer{}
	var promptServiceServer *server.PromptServiceServer = new(server.PromptServiceServer)

	client := db.InitializeDatabase(ctx)

	err = client.Ping(ctx, nil) // Ping the database to check if connection is successful
	if err != nil {
		panic(err)
	}

	defer db.DisconnectDatabase(ctx, client) // Disconnect from database when function ends

	answerServiceServer.SetDatabase(client)
	promptServiceServer.SetDatabase(client)

	// Function takes in pb.AnswerServiceServer, which is an interface.
	// Interfaces should be assigned pointers to structure implementations
	// in order to call value and pointer receiver methods.
	pb.RegisterAnswerServiceServer(grpcServer, answerServiceServer)
	pb.RegisterPromptServiceServer(grpcServer, promptServiceServer)

	fmt.Println("Server is running on ", port)
	grpcServer.Serve(lis) // Spawn a new goroutine to serve incoming requests
}
