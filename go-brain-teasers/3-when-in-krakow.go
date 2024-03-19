package main

import (
	"fmt"
)

func main() {
	// rune in Go is a data type that stores codes that represent Unicode characters.
	// the rune “ó” is taking two bytes
	city := "Kraków"
	fmt.Println(len(city))
}