package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// rune in Go is a data type that stores codes that represent Unicode characters.
	// the rune “ó” is taking two bytes
	city := "Kraków"
	fmt.Println(len(city))

	// To know the number of runes in a string
	fmt.Println(utf8.RuneCountInString(city))
}