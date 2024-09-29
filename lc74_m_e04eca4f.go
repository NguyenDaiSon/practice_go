package main

import (
	"fmt"
)

// //////////////////////////////////////////////
// LC: 74
// //////////////////////////////////////////////
func searchMatrix(m [][]int, target int) bool {
	rows := len(m)
	if rows == 0 {
		return false
	}

	cols := len(m[0])
	if cols == 0 {
		return false
	}

	left, right := 0, rows*cols-1
	for left <= right {
		mid := left + (right-left)/2
		val := m[mid/cols][mid%cols]
		if val == target {
			return true
		}

		if val < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1}}, 0))
}
