package main
import "fmt"
func sendValues(myIntChannel chan int){

  for i:=0; i<5; i++ {
    myIntChannel <- i //sending value 
  }
  close(myIntChannel)
}

func main() {
  myIntChannel := make(chan int)

  go sendValues(myIntChannel)

  for value := range myIntChannel {
    fmt.Println(value) //receiving value
  }
}