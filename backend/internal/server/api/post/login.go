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

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not parse json"))
		return
	}

	jwtToken, err := LoginInUser(user, r.UserAgent())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{
		"jwt": jwtToken,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	return

}

func LoginInUser(user models.User, userAgent string) (string, error) {
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

	updateRefreshToken(possibleUser, userAgent)

	return jwtToken, err
}

func updateRefreshToken(user *models.User, userAgent string) error {
	token, err := authentication.CreateRefreshToken(*user, userAgent)

	if err != nil {
		return err
	}

	db.UpdateUserRefreshToken(*user, *token)

	return nil
}
