package controllers

import "api/repo"

func GetAll() []repo.Book {
	return repo.GetBooks()
}
