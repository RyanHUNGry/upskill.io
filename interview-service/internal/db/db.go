package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitializeDatabase(ctx context.Context) *mongo.Client {
	URI := os.Getenv("INTERVIEW_SERVICE_MONGO_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))

	if err != nil {
		panic(err)
	}

	return client
}

func DisconnectDatabase(ctx context.Context, client *mongo.Client) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
