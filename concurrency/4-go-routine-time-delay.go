package main
import (
  "fmt"
  "time"
)

func WelcomeMessage(){
    fmt.Println("Welcome to Educative!")
}

func main() {
  go WelcomeMessage()
  go func(){  
    fmt.Println("Hello World!")
  }()

  time.Sleep(time.Millisecond*200)
}