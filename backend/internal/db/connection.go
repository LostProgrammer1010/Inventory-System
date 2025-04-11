package db

import (
	"context"
	"log"
	"os"

	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init() {

	connect()

	userCollection = client.Database("InventorySystem").Collection("Users")
	poleCollection = client.Database("InventorySystem").Collection("Poles")
	organizationCollection = client.Database("InventorySystem").Collection("Organization")

	createUniqueIndexes(client.Database("InventorySystem").Collection("Users"), []string{"email", "username"})
}

func connect() {
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

func createUniqueIndexes(collection *mongo.Collection, uniqueKeys []string) error {
	var indexes []mongo.IndexModel
	for _, key := range uniqueKeys {
		indexes = append(indexes, mongo.IndexModel{
			Keys:    bson.D{{Key: key, Value: 1}},
			Options: options.Index().SetUnique(true),
		})
	}

	_, err := collection.Indexes().CreateMany(context.TODO(), indexes)
	return err
}
