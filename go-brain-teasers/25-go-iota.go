package main

import (
	"fmt"
)

const (
	Read = 1 << iota
	Write
	Execute
)

func main() {
	fmt.Println(Execute)
}

/*
iota is an enumerated type.
It can be used inside a const declaration.
For each constant in the same group, iota grows by one.

Read 1 << "iota" -> 1 << 0 -> 1
Write 1 << 1
Execute 1 << 2
*/