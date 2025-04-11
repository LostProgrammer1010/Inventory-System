package post

import (
	"encoding/json"
	"net/http"

	"github.com/LostProgrammer1010/InventorySystem/internal/authentication"
	"github.com/LostProgrammer1010/InventorySystem/internal/db"
	"github.com/LostProgrammer1010/InventorySystem/internal/models"
)

func AddPole(w http.ResponseWriter, r *http.Request) {
	var pole models.Pole

	if checkMethod(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	oAuthToken := r.Header.Get("Authorization")
	authentication.VerifyJWTToken(oAuthToken)

	err := json.NewDecoder(r.Body).Decode(&pole)

	if err != nil {
		return
	}

	err = db.AddPole(pole)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully added"))
	return
}
