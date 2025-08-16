package server

import (
	"log"
	"net/http"
)

func StartServer() {

	RunHandler()
	err := http.ListenAndServe(":5555", nil)

	if err != nil {

		log.Fatalf("The server is down %v", err)
	}

}
