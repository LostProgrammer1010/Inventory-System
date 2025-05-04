package post

import (
	"encoding/json"
	"net/http"

	"github.com/LostProgrammer1010/InventorySystem/internal/authentication"
	"github.com/LostProgrammer1010/InventorySystem/internal/db"
	"github.com/LostProgrammer1010/InventorySystem/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddOrganization(w http.ResponseWriter, r *http.Request) {

	var organization models.Organization

	if checkMethod(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	oAuthToken := r.Header.Get("Authorization")

	claims, err := authentication.VerifyJWTToken(oAuthToken)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Not authorized to make this request"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&organization)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid organization"))
		return
	}

	addOrganizationID, err := db.AddOrgranization(organization)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to add organization to Database"))
		return
	}

	userID, err := primitive.ObjectIDFromHex(claims["UserID"].(string))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to get primitve object from jwt token"))
		return
	}

	user, err := db.GetUserById(userID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to get primitve object from jwt token"))
		return
	}

	newOrganization := models.OrganizationAuthorization{OrganizationID: addOrganizationID, Role: "Owner"}

	user.OrganizationAuthorization = append(user.OrganizationAuthorization, newOrganization)

	err = db.UpdateUser(*user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to update user in DB with organization"))
		return
	}

	NewAuthToken, err := authentication.CreateJWTAuthenticationToken(*user)

	err = json.NewEncoder(w).Encode(map[string]string{
		"jwt": NewAuthToken,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully Added"))
	return

}
