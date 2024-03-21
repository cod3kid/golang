package main

import (
	"fmt"
)

func main() {
	// nil is not a type but a reserved word. A variable initialized to nil must have a type
	n := nil
	fmt.Println(n)

	// This will work as variable n has a type
	// var n *int = nil
	// fmt.Println(n)
}
