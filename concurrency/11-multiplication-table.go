package main
import ( "fmt"
          "sync"
          "time")
          
func printTable(n int, wg *sync.WaitGroup) {
  for i := 1; i <= 12; i++ {
    fmt.Printf("%d x %d = %d\n", i, n, n*i)
    time.Sleep(50 * time.Millisecond)
  }
  wg.Done()
}

func main() {
  var wg sync.WaitGroup
  
  for number := 2; number <= 12; number++ {
    wg.Add(1)
    go printTable(number,&wg)
  }

  wg.Wait()
}