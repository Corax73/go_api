package server

import (
	"api/router"
	"fmt"
	"net/http"
)

func Main() {
	http.HandleFunc("/", router.Test)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
