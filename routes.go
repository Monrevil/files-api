package main

import (
	"fmt"
	"net/http"
)

func setupRoutes() {
	port := ":8080"
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(port, nil)
	fmt.Println("Listening on: localhost" + port)
}
