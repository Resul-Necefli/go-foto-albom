package handlers

import (
	"Resul-Necefli/go-foto-albom/model"
	"Resul-Necefli/go-foto-albom/storage"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func RunHandler() {
	http.HandleFunc("/photos", PhosotCollectionHandler)
	http.HandleFunc("/photo/", PhotoResursHandler)
}

func PhosotCollectionHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:

		GetPhotosHandler(w, r)
	case http.MethodPost:

		CreatPhotoHandler(w, r)
	default:

		log.Println("[PhosotCollectionHandler] method not allowed")
		http.Error(w, " method not allowed", http.StatusMethodNotAllowed)

	}

}

func PhotoResursHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		GetPhotoHandler(w, r)
	case http.MethodPut:

		UpdatePhotoHandler(w, r)

	case http.MethodPatch:

		UpdatePhotoHandler(w, r)

	case http.MethodDelete:

		DeletePhotoHandler(w, r)

	default:
		log.Println("[PhosotCollectionHandler] method not allowed")
		http.Error(w, " method not allowed", http.StatusMethodNotAllowed)

	}
}
func DeletePhotoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {

		log.Println("[DeletePhotoHandler] method not allowed ")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	param := strings.TrimPrefix(r.URL.Path, "/photo/")

	id, err := strconv.Atoi(param)

	if err != nil {

		log.Printf("[DeletePhotoHandler  Atio] %v ", err)
		http.Error(w, "server bad request", http.StatusBadRequest)
		return

	}

	photoObj, err := storage.GetByIDPhoto(id)
	if err != nil {

		log.Printf("DeletePhotoHandler : %v", err)
		http.Error(w, "server bad request", http.StatusBadRequest)
		return
	}

	storage.DeletePhoto(photoObj)

	w.WriteHeader(http.StatusOK)

}

func GetPhotosHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		log.Println("[GetPhotosHandler] method not allowed ")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(model.Photos)

	if err != nil {
		log.Printf("Conversion to json failed %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func GetPhotoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		log.Println("[GetPhotoHandler] method not allowed ")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	params := strings.TrimPrefix(r.URL.Path, "/photo/")

	ID, err := strconv.Atoi(params)

	if err != nil {
		log.Println(" strocnov could not convert ")
		http.Error(w, "server bad request", http.StatusBadRequest)
		return
	}

	photoObj, err := storage.GetPhoto(ID)

	if err != nil {
		log.Printf("failed to find photo by id %d: %v", ID, err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return

	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(photoObj); err != nil {

		log.Printf("Could not convert to json %v ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)

		return
	}

}

func UpdatePhotoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		log.Println("[GetPhotoHandler] method not allowed ")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var photoObj model.Photo

	err := json.NewDecoder(r.Body).Decode(&photoObj)

	if err != nil {

		log.Println("[UpdatePhotoHandler] Bad request")
		http.Error(w, " server bad request", http.StatusBadRequest)
		return
	}

	paramsID := strings.TrimPrefix(r.URL.Path, "/photo/")

	convertID, err := strconv.Atoi(paramsID)

	if err != nil {

		log.Printf("[ UpdatePhotoHandler] id not convert %v ", err)
		http.Error(w, "server bad request", http.StatusBadRequest)
		return

	}

	_, err = storage.GetByIDPhoto(convertID)

	if err != nil {

		log.Printf("[UpdatePhotoHandler] GetByIDPhoto  error  %v", err)
		http.Error(w, "information not found", http.StatusNotFound)
		return
	}

	photoObj.ID = convertID

	storage.UpdatePhoto(photoObj)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "photo updated successfully",
	})

}

func CreatPhotoHandler(w http.ResponseWriter, r *http.Request) {

	var obj model.Photo

	err := json.NewDecoder(r.Body).Decode(&obj)

	if err != nil {
		log.Printf("json format could not be decoded %v", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = storage.AddPhoto(obj)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	w.WriteHeader(http.StatusCreated)

}
