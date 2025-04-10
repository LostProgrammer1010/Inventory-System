package models

import (
	"time"
)

type RefreshToken struct {
	Token     string    `json:"token" bson:"token"`
	UserAgent string    `bson:"user_agent"`
	CreatedAt time.Time `bson:"created_at"`
	ExpiresAt time.Time `bson:"expires_at"`
}
