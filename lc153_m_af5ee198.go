package main

import (
	"fmt"
)

// ////////////////////////////////
// LC: 153
// ////////////////////////////////
func findMin(a []int) int {
	n := len(a)
	if n == 0 {
		return -1
	}

	left, right := 0, n-1
	index := 0

	for left <= right {
		mid := left + (right-left)/2
		if a[mid] <= a[n-1] {
			index = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return a[index]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4,5,6,7,0,1,2}))
	fmt.Println(findMin([]int{11,13,15,17}))
}
