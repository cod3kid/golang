package main

import (
"sync"
"fmt"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string){
defer wg.Done()
msg = s
}

func main(){
msg = "Hello World"

wg.Add(2)
go updateMessage("Hello, Universe")
go updateMessage("Hello, Cosmos")
wg.Wait()

fmt.Println(msg)
}