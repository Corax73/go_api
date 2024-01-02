package repo

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

type Book struct {
	Id     int
	Title  string
	Author string
}

type DataForCreatingBook struct {
	Title string
	Author string
	Slug string
	Description string
	Cover string
}

func getConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:@/library")
	if err == nil {
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
		return db
	}
	return nil
}

func GetBooks() []Book {
	books := []Book{}
	database = getConnect()
	if database != nil {
		rows, err := database.Query("select id, title, author from library.Books limit 10000")
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				book := Book{}
				err := rows.Scan(&book.Id, &book.Title, &book.Author)
				if err != nil {
					fmt.Println(err)
					continue
				}
				books = append(books, book)
			}
			return books
		}
	}
	return books
}

func CreatingBook(data DataForCreatingBook) bool {
	var resp bool
	database = getConnect()
	if database != nil {
		queryStr := "INSERT INTO library.Books(title, author, slug, description, cover) VALUES (" + "'" + data.Title + "'" + ", " + "'" + data.Author + "'" + ", " + "'" + data.Slug + "'" + ", " + "'" + data.Description + "'" + ", " + "'" + data.Cover + "'" + ")"
		rows, err := database.Query(queryStr)
		if err == nil {
			defer rows.Close()
			resp = true
		}
	}
	return resp
}
