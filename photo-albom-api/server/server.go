package server

import (
	"Resul-Necefli/go-foto-albom/handlers"
	"log"
	"net/http"
)

func StartServer() {

	handlers.RunHandler()
	err := http.ListenAndServe(":5555", nil)

	if err != nil {

		log.Fatalf("The server is down %v", err)
	}

}
