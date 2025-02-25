package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)

	go func() {

		for {
			select {
			case <-ctxWithCancel.Done():
				fmt.Println("Cancelled")
				return
			default:
				fmt.Println("Working")
				time.Sleep(time.Second * 1)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	cancel()
	fmt.Println("Canceling context...")
	time.Sleep(time.Second * 1)

}
