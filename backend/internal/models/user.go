package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                        primitive.ObjectID          `json:"_id,omitempty" bson:"_id,omitempty"`
	Username                  string                      `json:"username" bson:"username"`
	FirstName                 string                      `json:"first_name" bson:"first_name"`
	LastName                  string                      `json:"last_name" bson:"last_name"`
	Password                  string                      `json:"password" bson:"password"`
	Email                     string                      `json:"email" bson:"email"`
	OrganizationAuthorization []OrganizationAuthorization `json:"organization_auth" bson:"organization_auth"`
	RefreshToken              []RefreshToken              `bason:"refresh_token"`
}
