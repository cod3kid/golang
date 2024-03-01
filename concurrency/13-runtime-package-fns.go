package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("GOMAXPROCS is %d\n", runtime.GOMAXPROCS(2))
	fmt.Println(runtime.NumCPU())
}
