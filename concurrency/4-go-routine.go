package main
import "fmt"

func WelcomeMessage(){
    fmt.Println("Welcome Sufail!")
}

func main() {
  go WelcomeMessage()
  fmt.Println("Hello World!")
}