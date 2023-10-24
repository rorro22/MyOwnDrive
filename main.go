package main

import (
	"MyOwnDrive/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/upload", handlers.UploadHandler)
	http.HandleFunc("/files", handlers.FetchFilesHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
