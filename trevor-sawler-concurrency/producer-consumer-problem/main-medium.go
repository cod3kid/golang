package main

import (
	"fmt"
	"sync"
	"time"
)

// Producer generates numbers
func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the producer finishes
	for i := 1; i <= 5; i++ {
		fmt.Println("Produced:", i)
		ch <- i // Send the item to the channel
		time.Sleep(time.Millisecond * 500) // Simulate production time
	}
	close(ch) // Close the channel after production is complete
}

// Consumer processes numbers
func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the consumer finishes
	for item := range ch {
		fmt.Println("Consumed:", item)
		time.Sleep(time.Millisecond * 1000) // Simulate consumption time
	}
}

func main() {
	buffer := make(chan int, 2) // Shared buffer with capacity of 2
	var wg sync.WaitGroup       // Create a WaitGroup

	wg.Add(1) // Add 1 for the producer
	go producer(buffer, &wg)

	wg.Add(1) // Add 1 for the consumer
	go consumer(buffer, &wg)

	wg.Wait() // Wait for both producer and consumer to finish
	fmt.Println("All tasks completed!")
}
