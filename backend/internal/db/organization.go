package db

import (
	"context"

	"github.com/LostProgrammer1010/InventorySystem/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var organizationCollection *mongo.Collection

func AddOrgranization(organization models.Organization) (primitive.ObjectID, error) {

	result, err := organizationCollection.InsertOne(context.TODO(), organization)

	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil

}
