package main
import ( "fmt"
         "math/rand" 
         "time"
         "sync")

var wg sync.WaitGroup

func Race(channel, quit chan string, i int) {
  
  channel <- fmt.Sprintf("Car %d started!", i)  
    for{
      rand.Seed(time.Now().UnixNano())
      time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
      quit <- fmt.Sprintf("Car %d reached the finishing line!", i)
      fmt.Println(<-quit)
      wg.Done()
    }
}

func main() {

  channel := make(chan string)
  quit := make(chan string)
  wg.Add(1)
  for i:=0; i < 3; i++{
    go Race(channel,quit,i)
  }

  for{
    select{
      case raceUpdates := <-channel:
        fmt.Println(raceUpdates)
      case winnerAnnoucement := <-quit:
        fmt.Println(winnerAnnoucement)
        quit <- "You win!"
        wg.Wait()  
        return
      
    }
  }
}