package main

import (
	"log"
	"net/http"
)

func main(){
	mux := routes()

	log.Println("Starting Web Server on Port 8080")

	_ = http.ListenAndServe(":8080",mux)
}