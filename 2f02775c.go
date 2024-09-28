package main

import (
	"fmt"
)

func bubble_sort(a []int) []int {
	n := len(a)
	for i := n - 1; i >= 0; i-- {
		swapped := false
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return a
}

func main() {
	fmt.Println(bubble_sort([]int{2, 1, 3, 4}))
}
