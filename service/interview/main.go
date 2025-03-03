// Interview service driver

package main

import (
	"context"
	"interview/src/db"
	"interview/src/db/table"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	envName := ".env"
	envPath := filepath.Join(getWorkingDirectory(), envName)
	err := godotenv.Load(envPath)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initializeCassandra()
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

	// initialize tables if they do not exist
	table.InitializeTables(db.Session, db.Ctx)
}

func getWorkingDirectory() string {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal("Failed to get working directory")
	}

	return dir
}
