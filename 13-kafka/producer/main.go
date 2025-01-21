package main

import (
	"log"
	"net/http"
)

type Order struct {
	CustomerName string `json:"customer_name"`
	CoffeeType   string `json:"coffee_type"`
}

func main(){
	http.HandleFunc("/order",handleOrder)
	log.Fatal(http.ListenAndServe(":3000",nil))
}

func handleOrder(w http.ResponseWriter,req *http.Request){
	if req.Method != http.MethodPost{
		http.Error(w, "Invalid Request Method",http.StatusMethodNotAllowed)
		return
	}
}
