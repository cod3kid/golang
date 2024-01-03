package main

import (
"fmt"
"sync"
)


func printSomething(s string, wg *sync.WaitGroup){
	defer wg.Done()

	fmt.Println(s)
}
func main(){
	var wg sync.WaitGroup

	wg.Add(9)
	words := []string{
		"alpha",
		"beta",
		"gamma",
		"omega",
		"delta",
		"epsilon",
		"zeta",
		"eta",
		"theta",
	}	

	for i,x := range words{
		go printSomething(fmt.Sprintf("%d %s",i,x),&wg)
	}

	wg.Wait()

	wg.Add(1)
	printSomething("Second",&wg)
}