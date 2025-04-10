package post

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/InventorySystem/internal/authentication"
	"github.com/LostProgrammer1010/InventorySystem/internal/db"
	"github.com/LostProgrammer1010/InventorySystem/internal/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	w.Header().Set("Content-Type", "application/json")
	if checkMethod(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)

	jwtToken, err := LoginInUser(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(jwtToken)

}

func LoginInUser(user models.User) (string, error) {
	possibleUser, err := db.GetUserByUsername(user.Username)

	if err != nil {
		return "", fmt.Errorf("User was not found in db")
	}

	if !authentication.VerifyPassword(possibleUser.Password, user.Password) {
		return "", fmt.Errorf("Passwords did not match")
	}

	jwtToken, err := authentication.CreateJWTAuthenticationToken(*possibleUser)

	if err != nil {
		return "", err
	}

	return jwtToken, err
}
