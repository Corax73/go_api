package router

import (
	"api/controllers"
	"encoding/json"
	"net/http"
)

func HandlerBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	switch r.Method {
	case http.MethodGet:
		for _, book := range controllers.GetAll() {
			json.NewEncoder(w).Encode(book)
		}
	case http.MethodPost:
		if controllers.StoreBook(w, r) {
			json.NewEncoder(w).Encode("book created")
		}
	}
}
