package main

import (
	"fmt"
)

// //////////////////////////////////////
// LC: 69
// //////////////////////////////////////
func approximateSquareRoot(x int) int {
	if x <= 1 {
		return x
	}

	left, right := 1, x/2
	index := -1

	for left <= right {
		mid := left + (right-left)/2
		if mid == x/mid {
			index = mid
			break
		}

		if mid < x/mid {
			index = mid
			left = mid + 1
		} else {
			right = mid - 1
			index = right
		}
	}

	return index
}

func main() {
	for i := 0; i < 129; i++ {
		fmt.Printf("%d -> %d\n", i, approximateSquareRoot(i))
	}
}
