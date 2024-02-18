package main
import "fmt"
func sendValues(myIntChannel chan int){

  for i:=0; i<5; i++ {
    myIntChannel <- i //sending data through a channel 
  }

}

func main() {
  myIntChannel := make(chan int)

  go sendValues(myIntChannel)

  for i:=0; i<5; i++ {
    fmt.Println(<-myIntChannel) //receiving data through a channel
  }
}