// Driver for the interview service process.
package main

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"ryhung.upskill.io/internal/cassandra"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()

	// TODO: initialize Kafka queue

	// TODO: initialize Cassandra database
	fmt.Println("Initializing Cassandra database ⌛")
	db := cassandra.InitializeInterviewServiceDatabase(ctx)
	defer db.CloseInterviewServiceDatabase()

	// TODO: initialize gRPC API server

	for {

	}
}
