package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	client         *mongo.Client
	Users, Dailies *mongo.Collection
)

func Init() {
	opts := options.Client().ApplyURI(os.Getenv("DB_CONNECTION"))
	localClient, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = localClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}
	client = localClient
	Users = client.Database("daily-ai").Collection("users")
	Dailies = client.Database("daily-ai").Collection("dailies")
}

func Close() error {
	return client.Disconnect(context.Background())
}
