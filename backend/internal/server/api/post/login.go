package post

import (
	"encoding/json"
	"net/http"

	"github.com/LostProgrammer1010/InventorySystem/internal/authentication"
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

	jwtToken, err := authentication.LoginInUser(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
