package db

import (
	"context"

	"github.com/LostProgrammer1010/InventorySystem/internal/authentication"
	"github.com/LostProgrammer1010/InventorySystem/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var userCollection *mongo.Collection

// Adds the provided user to the db. Returns error if adding user failed
func AddUser(user models.User) (err error) {

	user.Password, err = authentication.HashPassword(user.Password)

	if err != nil {
		return err
	}

	_, err = userCollection.InsertOne(context.TODO(), user) // _ is the result object that was created

	if err != nil {
		return err
	}

	return nil

}

// Retrieves a user from the db and return the users if they were found
func GetUser(id primitive.ObjectID) (foundUser *models.User, err error) {
	filter := bson.M{"_id": id}

	err = userCollection.FindOne(context.TODO(), filter).Decode(&foundUser)

	if err != nil {
		return nil, err
	}

	return
}

// Retrives a user based on the username if any user was found
func GetUserByUsername(username string) (foundUser *models.User, err error) {
	filter := bson.M{"username": username}

	err = userCollection.FindOne(context.TODO(), filter).Decode(&foundUser)

	if err != nil {
		return nil, err
	}

	return
}
