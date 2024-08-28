// Driver for the interview service process.
package main

import (
	"context"

	"ryhung.upskill.io/internal/cassandra"
)

func main() {
	ctx := context.TODO()
	db := cassandra.InitializeInterviewServiceDatabase(ctx)
	defer db.CloseInterviewServiceDatabase()

	for {

	}
}
