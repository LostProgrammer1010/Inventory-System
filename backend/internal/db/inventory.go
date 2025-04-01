package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Item represents the structure of your inventory item
type Item struct {
	ID          string    `bson:"_id,omitempty"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	Quantity    int       `bson:"quantity"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

// GetInventoryCollection returns the inventory collection
func GetInventoryCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("inventory_system").Collection("inventory")
}

// CreateItem adds a new item to the inventory
func CreateItem(client *mongo.Client, item Item) error {
	collection := GetInventoryCollection(client)

	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	_, err := collection.InsertOne(context.TODO(), item)
	return err
}
