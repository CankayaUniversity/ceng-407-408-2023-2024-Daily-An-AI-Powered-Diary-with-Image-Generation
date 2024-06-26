package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	client                                           *mongo.Client
	Users, Dailies, ReportedDailies, UserPreferences *mongo.Collection
)

func Init() {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("DB_CONNECTION")).SetServerAPIOptions(serverAPI).SetConnectTimeout(120 * time.Second)

	localClient, err := mongo.Connect(context.Background(), opts)

	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = localClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}
	client = localClient
	UserPreferences = client.Database("daily").Collection("user-preferences")
	Users = client.Database("daily").Collection("users")
	Dailies = client.Database("daily").Collection("dailies")
	ReportedDailies = client.Database("daily").Collection("reportedDailies")
}

func Close() error {
	return client.Disconnect(context.Background())
}
