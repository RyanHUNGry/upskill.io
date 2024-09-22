// Driver for the interview service process.
package main

import (
	"context"
	"fmt"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"ryhung.upskill.io/internal/cassandra"
	api "ryhung.upskill.io/internal/grpc"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()

	// TODO: initialize Kafka queue

	// initialize Cassandra database
	go func(ctx context.Context) {
		fmt.Println("Initializing Cassandra database ⌛")
		db := cassandra.InitializeInterviewServiceDatabase(ctx)
		fmt.Println("Cassandra database initialized")
		defer db.CloseInterviewServiceDatabase()
	}(ctx)

	// initialize gRPC API server
	go func() {
		fmt.Println("Initializing gRPC API server ⌛")
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9999))
		if err != nil {
			panic(err)
		}
		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		api.RegisterInterviewServiceServer(grpcServer, api.CreateInterviewServiceServer())
		fmt.Println("gRPC API server initialized ✅")
		grpcServer.Serve(lis) // each gRPC request will be handled in a separate goroutine
		defer grpcServer.GracefulStop()
	}()

	// don't stop process
	select {}
}
