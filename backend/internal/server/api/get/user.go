package get

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/LostProgrammer1010/InventorySystem/internal/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path, "/user/")[1]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid User"))
		return
	}
	pID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to Parse ID into Primative Type"))
		return
	}

	foundUser, _ := db.GetUserById(pID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(foundUser)
	return
}
