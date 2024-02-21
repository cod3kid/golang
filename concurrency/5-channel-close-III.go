package main
import "fmt"
func sendValues(myIntChannel chan int){

  for i:=0; i<5; i++ {
    myIntChannel <- i //sending value 
  }
  
}

func main() {
  myIntChannel := make(chan int)
  go sendValues(myIntChannel)
  defer close(myIntChannel)

  for i:=0; i<5; i++ {
    fmt.Println(<-myIntChannel) //receiving value
  }
}