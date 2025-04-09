package post

import (
	"encoding/json"
	"net/http"

	"github.com/LostProgrammer1010/InventorySystem/internal/db"
	"github.com/LostProgrammer1010/InventorySystem/internal/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
	}

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = db.AddUser(newUser)

	return

}
