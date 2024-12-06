package main

import (
	"fmt"
	"time"
)

type Log struct {
	Message string
	time.Time
}

func main() {
	ts := time.Date(2009, 11, 10, 0, 0, 0, 0, time.UTC)
	log := Log{"Hello", ts}
	fmt.Printf("%v\n", log)
}

/*
%v prints all the fields but,
In the Log struct, there is a field with no name,
but only a type. This is called embedding,
and it means that the Log type has all the methods and fields that time.Time has.
*/