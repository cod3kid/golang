package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 6, 8, 20}
	fmt.Println(sum(numbers...))

	// fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}

func sum(v ...int) int {

	var total int

	for _, num := range v {
		total += num
	}

	return total
}
