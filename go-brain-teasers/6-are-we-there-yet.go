package main

import (
	"fmt"
	"time"
)

func main() {
	// timeout := 3

	var timeout time.Duration = 3 

	fmt.Printf("before ")
	// if timeout is an int, it'll be invalid operation: timeout * time.Millisecond (mismatched types int and time.Duration)
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	fmt.Println("after")
}