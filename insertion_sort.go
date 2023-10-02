package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := []int{4, 17, 28, 40, 98, -3, 0}
	fmt.Println("Unsorted ===== ", slice)
	insertionsort(slice)
	fmt.Println("Sorted ====== ", slice)
}

// insertionsort sorts the given array of integers in ascending order using the insertion sort algorithm.
//
// Parameters:
// - items: an array of integers to be sorted.
//
// Return type: None.
func insertionsort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}
