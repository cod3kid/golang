package main

import "fmt"

func deposit(balance *int,amount int){
    *balance += amount 
}

func withdraw(balance *int, amount int){
    *balance -= amount
}

func main() {
  
    balance := 100 
    
    go deposit(&balance,10) //depositing 10

    withdraw(&balance, 50) //withdrawing 50

    fmt.Println(balance) 


}


// go run --race filename.go, the --race flag is used to find data races
