package main

import (
	"fmt"
)

func firstNotSmaller(a []int, target int) int {
	n := len(a)
	left, right := 0, n-1
	index := -1
	for left <= right {
		mid := left + (right-left)/2
		if a[mid] >= target {
			index = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return index
}

func main() {
	fmt.Println(firstNotSmaller([]int{1, 3, 3, 3, 3, 3}, 3))
	fmt.Println(firstNotSmaller([]int{1, 2, 2, 2, 2, 2, 2, 3, 5, 8, 8, 10}, 2))
	fmt.Println(firstNotSmaller([]int{1, 3, 3, 5, 8, 8, 10}, 2))
	fmt.Println(firstNotSmaller([]int{0}, 0))
	fmt.Println(firstNotSmaller([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
	fmt.Println(firstNotSmaller([]int{1, 3, 3, 5, 8, 8, 10}, 0))
	fmt.Println(firstNotSmaller([]int{1, 3, 3, 5, 8, 8, 10}, 1))
	fmt.Println(firstNotSmaller([]int{2, 3, 5, 7, 11, 13, 17, 19}, 6))
}
