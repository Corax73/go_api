package server

import (
	"api/router"
	"fmt"
	"net/http"
)

func Main() {
	http.HandleFunc("/", router.GetAllBooks)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
