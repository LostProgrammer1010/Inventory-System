package authentication

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/LostProgrammer1010/InventorySystem/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

// Will need to be updated to something else in env to set the jwtKey values
var jwtKey = []byte("Very Secret Key")

func CreateJWTAuthenticationToken(user models.User) (string, error) {
	var authorizedOrganization []models.Organization

	for _, organization := range user.Organization {
		if organization.Role == "ADMIN" || organization.Role == "OWNER" {
			authorizedOrganization = append(authorizedOrganization, organization)
		}
	}

	claims := jwt.MapClaims{
		"UserID":           user.ID,
		"OrganizationAuth": authorizedOrganization,
		"ExpiresAt":        time.Now().Add(15 * time.Minute).Unix(),
		"GivenAt":          time.Now(),
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

func VerifyJWTToken(oAuth string) (any, error) {
	parts := strings.Split(oAuth, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, fmt.Errorf("Invalid Authorization header")
	}
	tokenString := parts[1]
	token, err := checkSignature(tokenString)

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims, nil
}

func checkSignature(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("Invalid Authorization Token")
	}

	return token, nil

}
