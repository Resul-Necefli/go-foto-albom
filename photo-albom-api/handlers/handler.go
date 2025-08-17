package handlers

import (
	"Resul-Necefli/go-foto-albom/jsonmanage"
	"Resul-Necefli/go-foto-albom/model"
	"Resul-Necefli/go-foto-albom/storage"
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
	http.HandleFunc("/photos/create", creatPhotoHandler)

}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Xos gelmisiniz\n")
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

	photoObj, err := storage.FindPhotoById(ID)

	if err != nil {

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return

	}

	w.Header().Set("content-type", "aplication/json")

	if err := json.NewEncoder(w).Encode(photoObj); err != nil {

		log.Printf("Could not convert to json %v ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)

		return
	}

}

func creatPhotoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var obj model.Photo

		err := json.NewDecoder(r.Body).Decode(&obj)

		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		err = storage.NewPhotoCreate(&obj)

		if err != nil {

			http.Error(w, err.Error(), http.StatusBadRequest)
			return

		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(obj)

		return
	}

	http.Error(w, "wrong method call", http.StatusMethodNotAllowed)

}
