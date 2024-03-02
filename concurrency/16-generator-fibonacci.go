package main

import "fmt"

func fibonacci(n int) chan int {
    mychannel := make(chan int)
    go func() {
        k := 0
        for i, j := 0, 1; k < n ; k++ {
            mychannel <- i
            i, j = i+j,i
            
        }
            close(mychannel)
    }()
    return mychannel
}

func main() {
  
    for i := range fibonacci(10) {
        fmt.Println(i)
    }
}