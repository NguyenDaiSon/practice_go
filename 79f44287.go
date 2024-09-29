package main

import (
	"fmt"
)

func findBoundary(a []bool) int {
	left, right := 0, len(a)-1
	for left <= right {
		mid := left + (right-left)/2
		if !a[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left >= len(a) {
		left = -1
	}
	return left
}

func main() {
	fmt.Println(findBoundary([]bool{true, true, true}))
	fmt.Println(findBoundary([]bool{false, false, true, true, true}))
	fmt.Println(findBoundary([]bool{false, false, false}))
}
