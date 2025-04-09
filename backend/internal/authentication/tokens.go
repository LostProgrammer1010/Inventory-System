package authentication

import (
	"github.com/LostProgrammer1010/InventorySystem/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("Very Secret Key")

func createJWTToken(user models.User) {
	claims := jwt.MapClaims{}
}
