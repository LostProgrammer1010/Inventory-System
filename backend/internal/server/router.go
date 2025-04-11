package server

import (
	"net/http"

	"github.com/LostProgrammer1010/InventorySystem/internal/server/api/get"
	"github.com/LostProgrammer1010/InventorySystem/internal/server/api/post"
)

func createRouter() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	router.HandleFunc("/user/add", post.AddUser)
	router.HandleFunc("/user/", get.GetUserById)
	router.HandleFunc("/user/login", post.Login)
	router.HandleFunc("/pole/add", post.AddPole)
	router.HandleFunc("/organization/add", post.AddOrganization)

	return router
}
