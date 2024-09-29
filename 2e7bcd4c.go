package main

import (
	"fmt"
)

func binary_search(a []int, target int) int {
	left, right := 0, len(a)-1
	for left <= right {
		middle := left + (right-left)/2
		if a[middle] == target {
			return middle
		}

		if a[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}

func main() {
	fmt.Println(binary_search([]int{1}, 2))
	fmt.Println(binary_search([]int{}, 0))
	for i := 0; i < 5; i++ {
		fmt.Println(binary_search([]int{1, 2, 3, 4}, i))
	}
}
