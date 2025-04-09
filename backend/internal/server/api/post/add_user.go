package post

import (
	"encoding/json"
	"github.com/LostProgrammer1010/InventorySystem/internal/db"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser db.User
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
	}

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	db.CreateUser(newUser)

}
