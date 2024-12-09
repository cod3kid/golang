package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var wg sync.WaitGroup

	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
		}()

	}
	wg.Wait()
	fmt.Println(count)
}


/*
It'll print a random number under 1000000 due to race condition
*/