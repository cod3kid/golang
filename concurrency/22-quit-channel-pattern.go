package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Race(channel, quit chan string, i int) {

	channel <- fmt.Sprintf("Car %d has started", i)
	for {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		quit <- fmt.Sprintf("Car %d reached the finishing line!", i)
	}
}

func main() {
	channel := make(chan string)
	quit := make(chan string)

	for i := 0; i < 3; i++ {
		go Race(channel, quit, i)
	}

	for {
		select {
		case raceUpdates := <-channel:
			fmt.Println(raceUpdates)
		case winner := <-quit:
			fmt.Println(winner)
			return
		}
	}

}
