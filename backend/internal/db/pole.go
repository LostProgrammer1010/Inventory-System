package db

import (
	"context"

	"github.com/LostProgrammer1010/InventorySystem/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var poleCollection *mongo.Collection

func AddPole(pole models.Pole) error {

	_, err := poleCollection.InsertOne(context.TODO(), pole)

	if err != nil {
		return err
	}

	return nil

}
