package main
import "fmt"

type Money struct{
  amount int
  year int
}

func sendMoney(parent chan Money){

  for i:=0; i<=18; i++ {
    parent <- Money{5000,i}  
  }
  close(parent)
}

func main() {
  money := make(chan Money)

  go sendMoney(money)

  for kidMoney:= range money {
    fmt.Printf("Money received by kid in year %d : %d\n", kidMoney.year, kidMoney.amount) 
  }
}