package db

import (
	"context"

	"github.com/LostProgrammer1010/InventorySystem/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var organizationCollection *mongo.Collection

func AddOrgranization(organization models.Organization) error {

	_, err := organizationCollection.InsertOne(context.TODO(), organization)

	if err != nil {
		return err
	}

	return nil

}
