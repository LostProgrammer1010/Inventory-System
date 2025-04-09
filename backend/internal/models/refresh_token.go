package models

type RefreshToken struct {
	Token     string `bson:"token"`
	UserAgent string `bson:"user_agent"`
	ExpiresAt int64  `bson:"expires_at"`
}
