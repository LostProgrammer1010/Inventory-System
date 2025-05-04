package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Organization struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
}

type OrganizationAuthorization struct {
	OrganizationID primitive.ObjectID `json:"organization_id" bson:"organization_id"`
	Role           string             `json:"Role" bson:"Role"`
}
