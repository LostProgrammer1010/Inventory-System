package post

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/LostProgrammer1010/InventorySystem/internal/authentication"
	"github.com/LostProgrammer1010/InventorySystem/internal/db"
	"github.com/LostProgrammer1010/InventorySystem/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddPole(w http.ResponseWriter, r *http.Request) {
	var pole models.Pole

	if checkMethod(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	oAuthToken := r.Header.Get("Authorization")
	claims, err := authentication.VerifyJWTToken(oAuthToken)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You do not have permission to make this request"))
		return
	}

	authorizedOrganization := claims["OrganizationAuth"].([]any)

	if len(authorizedOrganization) == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You do not have permission to make this request"))
		return
	}

	organization_requested, err := primitive.ObjectIDFromHex(strings.Split(r.URL.Path, "/")[1])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to create primitve object"))
		return
	}

	authorized := false

	for _, organization := range authorizedOrganization {
		var auth models.OrganizationAuthorization
		encoded, _ := json.Marshal(organization)
		json.Unmarshal(encoded, &auth)
		if auth.OrganizationID == organization_requested {
			authorized = true
			break
		}
	}

	if !authorized {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You do are not authorized to make this request not owner or admin of organization"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&pole)

	if err != nil {
		return
	}

	pole.Organization = organization_requested

	err = db.AddPole(pole)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully added"))
	return
}
