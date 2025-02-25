// You can edit this code!
// Click here and start typing.
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctxWithTimeout,cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	done := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Api called")
	case <-ctxWithTimeout.Done():
		fmt.Println("Time limit exceeded")
	}

}
