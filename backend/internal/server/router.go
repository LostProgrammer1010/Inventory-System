package server

import (
	"github.com/LostProgrammer1010/InventorySystem/internal/server/api/post"
	"net/http"
)

func createRouter() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	router.HandleFunc("/user/add", post.CreateUser)

	return router
}
