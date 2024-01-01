package router

import (
	"api/controllers"
	"encoding/json"
	"net/http"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	for _, book := range controllers.GetAll() {
		json.NewEncoder(w).Encode(book)
	}
}
