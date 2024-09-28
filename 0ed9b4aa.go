package main

import (
	"fmt"
)

func select_sort(a []int) []int {
	n := len(a)
	for i := 0; i < n-1; i++ {
		m := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[m] {
				m = j
			}
		}

		if m != i {
			a[m], a[i] = a[i], a[m]
		}
	}
	return a
}

func main() {
	fmt.Println(select_sort([]int{1, 3, 2, 4}))
}
