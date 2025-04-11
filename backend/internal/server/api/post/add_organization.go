package post

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/InventorySystem/internal/authentication"
	"github.com/LostProgrammer1010/InventorySystem/internal/db"
	"github.com/LostProgrammer1010/InventorySystem/internal/models"
)

func AddOrganization(w http.ResponseWriter, r *http.Request) {

	var organization models.Organization

	if checkMethod(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	oAuthToken := r.Header.Get("Authorization")

	claims, err := authentication.VerifyJWTToken(oAuthToken)
	fmt.Println(claims)

	err = json.NewDecoder(r.Body).Decode(&organization)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.AddOrgranization(organization)

	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully Added"))

}
