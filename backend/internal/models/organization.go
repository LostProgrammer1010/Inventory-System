package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Organization struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
	Role string             `json:"role" bson:"role"`
}
