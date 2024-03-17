package main

import (
	"fmt"
)

func main() {
	var m map[string]int
	fmt.Println(m["Hello"])


	// Use the comma (,ok) paradigm to know if a key exist in a map
	_,ok:= m["Mello"]
	fmt.Println(ok)
}



/*
 Some operation on map is nil safe
 If the key is on the map, return its value.
 If the key is not on the map, return the zero value for the value type.

 Zero values for different types
Boolean: false.
Integer : 0.
Floating : 0.0.
String:""
Interfaces, slices, channels, maps, pointers, and functions: nil.
*/
