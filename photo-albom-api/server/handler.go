package server

import (
	"Resul-Necefli/go-foto-albom/jsonmanage"
	"Resul-Necefli/go-foto-albom/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func RunHandler() {

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/photos", photosHandler)
	http.HandleFunc("/photos/", photosQueryHanadler)

}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Xos gelmisiniz!")
}

func photosHandler(w http.ResponseWriter, r *http.Request) {

	dataByte, err := jsonmanage.Jsonconversion()

	if err != nil {
		log.Printf("Conversion to json failed %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)

	}

	w.Header().Set("content-type", "aplication/json")
	w.Write(dataByte)
}

func photosQueryHanadler(w http.ResponseWriter, r *http.Request) {

	params := strings.TrimPrefix(r.URL.String(), "/photos/")

	ID, err := strconv.Atoi(params)

	if err != nil {
		http.Error(w, "server bad request", http.StatusBadRequest)
		return
	}

	photoObje, err := model.FindPhotoById(ID)

	if err != nil {

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return

	}

	w.Header().Set("content-type", "aplication/json")

	if err := json.NewEncoder(w).Encode(photoObje); err != nil {

		log.Printf("Could not convert to json %v ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)

		return
	}

}
