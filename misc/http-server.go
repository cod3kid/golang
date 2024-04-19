package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/hello",func (w http.ResponseWriter,req *http.Request){
		fmt.Fprintln(w,"Hello Meow Meow")
	})

	http.ListenAndServe(":4000",nil)
}