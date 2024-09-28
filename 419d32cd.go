package main

import (
	"fmt"
)

func merge_sort(a []int) []int {
	n := len(a)
	if n <= 1 {
		return a
	}

	l, r := partition(a)
	l = merge_sort(l)
	r = merge_sort(r)
	a = merge(l, r)
	return a
}

func partition(a []int) ([]int, []int) {
	n := len(a)
	if n <= 1 {
		return a, nil
	}
	return a[:n/2], a[n/2:]
}

func merge(l, r []int) []int {
	a := make([]int, 0)
	i, j := 0, 0
	nl, nr := len(l), len(r)

	for i < nl && j < nr {
		if l[i] < r[j] {
			a = append(a, l[i])
			i += 1
		} else {
			a = append(a, r[j])
			j += 1
		}
	}

	for i < nl {
		a = append(a, l[i])
		i += 1
	}

	for j < nr {
		a = append(a, r[j])
		j += 1
	}

	return a
}

func main() {
	fmt.Println(merge_sort([]int{2, 1, 3, 4}))
}
