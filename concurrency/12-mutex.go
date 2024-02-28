package main
import (
  "fmt"
  "sync"
  "time")

func main() {
  var myMutex sync.Mutex
   myMutex.Lock()
   go func() {
    myMutex.Lock()
    fmt.Println("I am in the goroutine")
    myMutex.Unlock()
  }()
  fmt.Println("I am in main routine")
  myMutex.Unlock()
  time.Sleep(time.Second*1)
  
}