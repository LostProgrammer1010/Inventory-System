package db

import (
	"context"
	"fmt"
	"github.com/LostProgrammer1010/InventorySystem/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type RefreshToken struct {
	Token     string `bson:"token"`
	UserAgent string `bson:"user_agent"`
	ExpiresAt int64  `bson:"expires_at"`
}

func CreateUser(user models.User) {

	usersCollection := client.Database("InventorySystem").Collection("Users")

	result, err := usersCollection.InsertOne(context.TODO(), user)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

}

func GetUser(id primitive.ObjectID) (foundUser *User, err error) {
	usersCollection := client.Database("InventorySystem").Collection("Users")
	filter := bson.M{"_id": id}

	err = usersCollection.FindOne(context.TODO(), filter).Decode(&foundUser)

	if err != nil {
		return nil, err
	}

	return

}
