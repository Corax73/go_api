package controllers

import (
	"api/repo"
	"encoding/json"
	"net/http"
)

func GetAll() []repo.Book {
	return repo.GetBooks()
}

func StoreBook(w http.ResponseWriter, r *http.Request) bool {
	var NewBook repo.DataForCreatingBook
	var resp bool
	err := json.NewDecoder(r.Body).Decode(&NewBook)
	if err == nil {
		if repo.CreatingBook(NewBook) {
			resp = true
		}
	}
	return resp
}