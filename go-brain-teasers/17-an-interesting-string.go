package main

import (
	"fmt"
)

func main() {
	i := 169
	s := string(i)
	fmt.Println(s)
}


/*
Go will consider 169 as a rune, that's why it return copyright symbol
*/