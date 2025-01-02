package main

import (
	"fmt"
	"time"
)

func calculateNextInt(prev int) int {
	time.Sleep(1 * time.Second)
	return prev + 1
}

type producer struct {
    data chan int
    quit chan chan error
}
 
func (p *producer) Close() error {
    ch := make(chan error)
    p.quit <- ch
    return <-ch
}
 
func main() {
    prod := &producer{
        data: make(chan int),
        quit: make(chan chan error),
    }
 
    // producer
    go func() {
        var i = 0
        for {
            i = calculateNextInt(i)
            select {
            case prod.data <- i:
            case ch := <-prod.quit:
                close(prod.data)
                // If the producer had an error while shutting down,
                // we could write the error to the ch channel here.
                close(ch)
                return
            }
        }
    }()
 
    // consumer
    for i := range prod.data {
        fmt.Printf("Producer #%v\n", i)
        if i >= 5 {
            err := prod.Close()
            if err != nil {
                // cannot happen in this example
                fmt.Printf("unexpected error: %v\n", err)
            }
        }
    }
}