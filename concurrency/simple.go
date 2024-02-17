package main

import (
	"fmt"
	"runtime"
	"sync"
)

func valueSet3(goChannel chan<- int) {
   defer wg.Done()
	for i := 10; i < 15; i++ {
		goChannel <- i
	}
}

func valueSet2(goChannel chan<- int) {
	defer wg.Done()
	for i := 5; i < 10; i++ {
		goChannel <- i
	}
}

func valueSet1(goChannel chan<- int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		goChannel <- i
	}
}

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(4)
	wg.Add(3)
	goChannel := make(chan int)
	go valueSet1(goChannel)
	go valueSet3(goChannel)
	go valueSet2(goChannel)

	for i := 0; i < 15; i++ {
		fmt.Println(<-goChannel)
	}

	wg.Wait()
}
