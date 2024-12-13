package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	<-ch
	close(ch)
	a := <-ch
	b := <-ch
	fmt.Println(a, b)
}


/*
2 is the length of the buffered channel
Any data sent after will be 0

How can we know if a value we get from a channel is a value that was there,
or a zero value since the channel was closed?vWe use the comma (,ok) paradigm to check.

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	ch <- 1
	a, ok := <-ch
	fmt.Println(a, ok) // 1 true
	close(ch)
	b, ok := <-ch
	fmt.Println(b, ok) // 0 false
}
*/