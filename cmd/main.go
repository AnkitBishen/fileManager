package main

import (
	"fmt"
	"net/http"

	getFile "github.com/AnkitBishen/fileManagerApp/internal/handlers"
)

func main() {
	fmt.Println("Hello, World!")
	// get config

	// manage routes
	router := http.NewServeMux()
	router.HandleFunc("GET /api/getlist", getFile.List())

	// start server
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
