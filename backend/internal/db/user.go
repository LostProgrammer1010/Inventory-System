package db

import (
	"context"
	"fmt"
	"github.com/LostProgrammer1010/InventorySystem/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Adds the provided user to the db. Returns error if adding user failed
func AddUser(user models.User) error {

	usersCollection := client.Database("InventorySystem").Collection("Users")

	user.Password =  

	_, err := usersCollection.InsertOne(context.TODO(), user) // _ is the result object that was created

	if err != nil {
		return err
	}

	return nil

}

// Retrieves a user from the db and return the users if they were found
func GetUser(id primitive.ObjectID) (foundUser *models.User, err error) {
	usersCollection := client.Database("InventorySystem").Collection("Users")
	filter := bson.M{"_id": id}

	err = usersCollection.FindOne(context.TODO(), filter).Decode(&foundUser)

	if err != nil {
		return nil, err
	}

	return
}

// Retrives a user based on the username if any user was found
func GetUserByUsername(username string) (foundUser *models.User, err error) {
	userCollection := client.Database("InventorySystem").Collection("Users")
	filter := bson.M{"username": username}

	err = userCollection.FindOne(context.TODO, filter).Decode(&foundUser)

	if err != nil {
		return nil, err 
	}

	return
}

