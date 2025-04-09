package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Username     string             `json:"username" bson:"username"`
	Password     string             `json:"password" bson:"password"`
	Email        string             `json:"email" bson:"email"`
	Role         string             `json:"role" bson:"role"`
	Organization string             `json:"organization" bson:"organization"`
	refreshToken RefreshToken       `bason:"refreshtoken"`
}
