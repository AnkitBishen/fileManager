package main

import (
	"fmt"
	"net/http"

	"github.com/AnkitBishen/fileManagerApp/internal/crosMid"
	getFile "github.com/AnkitBishen/fileManagerApp/internal/handlers"
)

func main() {
	fmt.Println("Hello, World!")
	// get config

	// manage routes
	router := http.NewServeMux()
	router.HandleFunc("POST /api/getlist", getFile.List())

	crosMux := crosMid.CorsMiddleware(router)

	// start server
	err := http.ListenAndServe(":8080", crosMux)
	if err != nil {
		fmt.Println(err.Error())
	}
}
