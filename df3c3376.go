package main

import (
	"fmt"
)

func quick_sort(a []int) []int {
	// Call the recursive helper function to start sorting
	// from the beginning (index 0) to the end (length of the array).
	quick_sort_impl(a, 0, len(a))
	return a
}

func quick_sort_impl(a []int, start, end int) {
	// Base case: If the subarray has one or no elements, it's already sorted.
	if end-start <= 1 {
		return
	}

	// Step 1: Select the pivot
	// The pivot is chosen as the last element of the current subarray.
	pivot := a[end-1]

	// Step 2: Initialize two pointers
	// i points to the start, j starts at the element before the pivot.
	i, j := start, end-1

	// Step 3: Partitioning the array
	// Move i to the right and j to the left, swapping elements
	// that are on the wrong side of the pivot.
	for i < j {
		// Move i right until we find an element that is greater or equal to the pivot.
		for a[i] < pivot && i < j {
			i += 1
		}

		// Move j left until we find an element that is smaller than the pivot.
		for a[j] >= pivot && i < j {
			j -= 1
		}

		// Swap elements at i and j since they are on the wrong sides.
		a[i], a[j] = a[j], a[i]
	}

	// Step 4: Place the pivot in its correct position.
	// i now points to the first element greater than or equal to the pivot,
	// so swap it with the pivot to place the pivot in its sorted position.
	a[i], a[end-1] = a[end-1], a[i]

	// Step 5: Recursively apply quicksort to the left and right partitions.
	// Recursively sort the subarray left of the pivot.
	quick_sort_impl(a, start, i)
	// Recursively sort the subarray right of the pivot.
	quick_sort_impl(a, i+1, end)
}

func main() {
	// Example usage: Sort an array and print the result.
	fmt.Println(quick_sort([]int{2, 1, 3, 4}))
}
