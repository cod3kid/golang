package main

import (
	"fmt"
	"net/http"
	"time"
)


func cancellableHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			
	fmt.Println("Started processing request")

	select {
	case <-time.After(5 * time.Second): 
		fmt.Fprintln(w, "Finished work successfully")
		fmt.Println("Finished processing request")
	case <-ctx.Done():
		fmt.Println("Request canceled by the client:", ctx.Err())
		return
	}
}

func main(){

	http.HandleFunc("/cancellable",cancellableHandler)

	fmt.Println("Server listening to port 3000")
	err:=http.ListenAndServe(":3000",nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}

}