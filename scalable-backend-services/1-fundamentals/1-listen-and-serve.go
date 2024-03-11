package main

import (
	"log"
	"net/http"
	"io"
)

func handle(w http.ResponseWriter, r * http.Request){
	io.WriteString(w, "Hello World!\n")
}

func main() {
	http.HandleFunc("/",handle)
	log.Println("Hello")
	log.Fatal(http.ListenAndServe(":8000",nil))
}
