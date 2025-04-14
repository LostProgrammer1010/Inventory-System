package server

import (
	"github.com/LostProgrammer1010/InventorySystem/internal/server/api/get"
	"github.com/LostProgrammer1010/InventorySystem/internal/server/api/post"
	"github.com/gorilla/mux"
	"net/http"
)

func createRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	router.HandleFunc("/user/add", post.AddUser)
	router.HandleFunc("/user/", get.GetUserById)
	router.HandleFunc("/user/login", post.Login)
	router.HandleFunc("/{organization}/pole/add", post.AddPole) // organization/pole/add
	router.HandleFunc("/organization/add", post.AddOrganization)

	return router
}
