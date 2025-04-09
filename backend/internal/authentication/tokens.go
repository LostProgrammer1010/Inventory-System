package authentication

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/LostProgrammer1010/InventorySystem/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// Will need to be updated to something else in env to set the jwtKey values
var jwtKey = []byte("Very Secret Key")

func CreateJWTAuthenticationToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"userID":    user.ID,
		"expiresAt": time.Now().Add(15 * time.Minute).Unix(),
		"GivenAt":   time.Now(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)

}

func CreateRefreshToken(user models.User, broswerAgent string) (*models.RefreshToken, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return nil, err
	}

	refreshToken := models.RefreshToken{
		Token:     hex.EncodeToString(token),
		UserAgent: broswerAgent,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().AddDate(0, 0, 30),
	}

	return &refreshToken, nil
}
