package main

import "fmt"

func main() {

	arr := []int{1, 2, 3}
	result := [][]int{}

	var backtrack func(int, []int)

	backtrack = func(idx int, current []int) {
		if idx == len(arr) {
			currentCopy := make([]int, len(current))
			copy(currentCopy, current)
			result = append(result, currentCopy)
			return
		}

		backtrack(idx+1, current)

		current = append(current, arr[idx])
		backtrack(idx+1, current)

	}

	backtrack(0, []int{})

	fmt.Println(result)
}
