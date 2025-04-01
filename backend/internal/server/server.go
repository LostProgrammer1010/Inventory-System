package server

import (
	"fmt"
	"net/http"
)

func Start() {

	router := createRouter()
	address := "localhost"
	port := "8080"

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", address, port),
		Handler: router,
	}

	server.ListenAndServe()

}
