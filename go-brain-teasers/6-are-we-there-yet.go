package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := 3
	fmt.Printf("before ")
	//  invalid operation: timeout * time.Millisecond (mismatched types int and time.Duration)
	time.Sleep(timeout * time.Millisecond)
	fmt.Println("after")
}