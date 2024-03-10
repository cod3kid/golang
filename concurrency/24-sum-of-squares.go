package main

import "fmt"

func calcSum(myChan, quitChan chan int) {
	y := 1
	for {
		select {
		case myChan <- (y * y):
			y++
		case <-quitChan:
			return
		}
	}
}
func main() {

	myChannel := make(chan int)
	quitChannel := make(chan int)
	sum := 0

	go func() {
		for i := 1; i <= 5; i++ {
			sum += <-myChannel
		}
		fmt.Println(sum)
		quitChannel <- 0
	}()

 	calcSum(myChannel, quitChannel)
}

