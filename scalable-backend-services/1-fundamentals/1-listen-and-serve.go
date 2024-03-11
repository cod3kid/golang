package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println("Request path:", r.URL.Path)
	io.WriteString(w, "Hello students!\n")
}

func main() {
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/", handle)
	log.Printf("Starting server on localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}