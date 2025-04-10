package post

import (
	"encoding/json"
	"net/http"

	"github.com/LostProgrammer1010/InventorySystem/internal/db"
	"github.com/LostProgrammer1010/InventorySystem/internal/models"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invaid Method"))
		return
	}

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if !checkValidUserInput(newUser) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User missing either Username, Password, or Email"))
		return
	}

	err = db.AddUser(newUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("User was successfully added to database"))

	return

}

func checkValidUserInput(newUser models.User) bool {

	if len(newUser.Username) == 0 {
		return false
	}
	if len(newUser.Password) == 0 {
		return false
	}
	if len(newUser.Email) == 0 {
		return false
	}
	return true
}
