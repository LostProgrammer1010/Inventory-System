package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pole struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Brand        string             `json:"brand" bson:"brand"`
	Weight       int                `json:"weight" bson:"weight"`
	Length       Length             `json:"length" bson:"length"`
	Flex         float32            `json:"flex" bson:"flex"`
	Organization primitive.ObjectID `json:"organization" bson:"organization"`
	Renter       primitive.ObjectID `json:"renter" bson:"renter"`
}

type Length struct {
	Inches float32 `bson:"inches"`
	Feet   float32 `bson:"feet"`
	Meters float32 `bson:"meters"`
}
