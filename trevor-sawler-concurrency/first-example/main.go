package main

import (
"fmt"
"time"
)


func printSomething(s string){
	fmt.Println(s)
}
func main(){
	go printSomething("First")

	time.Sleep(1*time.Second)
	printSomething("Second")
}