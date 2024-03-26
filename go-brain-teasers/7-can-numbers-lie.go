package main

import (
	"fmt"
	"math"
)

func main() {
	//  we sacrifice accuracy for speed. It's the same JavaScript, Python etc...
	n := 1.1
	fmt.Println(n * n)

	//  The NaN value does not equal any number, including itself. It's the same in JavaScript
	fmt.Println(math.NaN() == math.NaN())
}