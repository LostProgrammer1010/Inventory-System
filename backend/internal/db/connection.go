package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Connect() {

	CONNECTION_STRING := os.Getenv("MONGODB_CONNECTION_STRING")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(CONNECTION_STRING).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal("Server Failed to Connect to MONGODB", err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatal("Server Failed to disconnect from MONGODB", err)
		}
	}()
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		log.Fatal("Server Failed to Ping CLUSTER", err)
	}
}
