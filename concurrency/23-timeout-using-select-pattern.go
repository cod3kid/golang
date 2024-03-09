package main
import ( "fmt"
         "time"
         "math/rand")


func main() {
  dynamite := make(chan string)
 
  go func(){
    rand.Seed(time.Now().UnixNano())
    time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
    dynamite <- "Dynamite Diffused!"
  }()
  
  timeout := time.After(time.Duration(rand.Intn(500)) * time.Millisecond) 
  for {
    select
    { case s := <-dynamite:
        fmt.Println(s)
        return
      case <-timeout:
        fmt.Println("Dynamite Explodes!") 
        return       
    }
  }
}