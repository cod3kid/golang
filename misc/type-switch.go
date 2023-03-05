package main

import "fmt"

func typeSwitch(i interface{}) {

	switch i.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Others")
	}
}

func main() {

	typeSwitch(2)
	typeSwitch(2.7)
	typeSwitch("hello")
}
