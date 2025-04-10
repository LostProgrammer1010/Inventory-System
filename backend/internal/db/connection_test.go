package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func testconnect() {
	var err error
	CONNECTION_STRING := os.Getenv("MONGODB_CONNECTION_STRING")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(CONNECTION_STRING).SetServerAPIOptions(serverAPI)
	client, err = mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal("Server Failed to Connect to MONGODB", err)
	}

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{bson.E{Key: "ping", Value: 1}}).Err(); err != nil {
		log.Fatal("Server Failed to Ping CLUSTER", err)
	}

	fmt.Println("Connected to MongoDB cluster Successfully")
}
