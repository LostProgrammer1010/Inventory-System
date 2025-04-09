package authentication

import (
	"fmt"

	"github.com/LostProgrammer1010/InventorySystem/internal/db"
	"github.com/LostProgrammer1010/InventorySystem/internal/models"
)

func LoginInUser(user models.User) (string, error) {
	possibleUser, err := db.GetUserByUsername(user.Username)

	if err != nil {
		return "", fmt.Errorf("User was not found in db")
	}

	if !VerifePassword(possibleUser.Password, user.Password) {
		return "", nil
	}

	jwtToken, err := CreateJWTAuthenticationToken(*possibleUser)

	if err != nil {
		return "", err
	}

	return jwtToken, err
}
