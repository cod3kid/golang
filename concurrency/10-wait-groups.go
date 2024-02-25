package main

import (
	"fmt"
	"sync"
)

func WelcomeMessage() {
	fmt.Println("Welcome to Educative!")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		WelcomeMessage()
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello World!")
		wg.Done()
	}()

	wg.Wait()
}
