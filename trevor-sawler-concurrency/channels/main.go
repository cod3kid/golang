package main

import (
	"fmt"
	"strings"
)

// shout has two parameters: a receive only chan ping, and a send only chan pong.
// Note the use of <- in function signature. It simply takes whatever
// string it gets from the ping channel,  converts it to uppercase and
// appends a few exclamation marks, and then sends the transformed text to the pong channel.
func shout(ping <-chan string, pong chan<- string) {
	for {
		// read from the ping channel. Note that the GoRoutine waits here -- it blocks until
		// something is received on this channel.
		s := <-ping

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	// create two channels. Ping is what we send to, and pong is what comes back.
	ping := make(chan string)
	pong := make(chan string)

	// start a goroutine
	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")

	for {
		// print a prompt
		fmt.Print("-> ")

		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			// jump out of for loop
			break
		}

		// send userInput to "ping" channel
		ping <- userInput

		// wait for a response from the pong channel. Again, program
		// blocks (pauses) until it receives something from 
		// that channel.
		response := <-pong

		// print the response to the console.
		fmt.Println("Response:", response)
	}

	fmt.Println("All done. Closing channels.")

	// close the channels
	close(ping)
	close(pong)
}
