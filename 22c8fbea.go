package main

import (
	"fmt"
)

func insert_sort(a []int) []int {
	n := len(a)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			a[j-1], a[j] = a[j], a[j-1]
		}
	}
	return a
}

func main() {
	fmt.Println(insert_sort([]int{1, 3, 2, 4}))
}
