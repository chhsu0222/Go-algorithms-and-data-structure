package main

import "fmt"

func main() {
	xi := []int{1, 1, 1, 1, 1, 2}
	count := countUniqueValues(xi)
	fmt.Println(count) // 2

	xi = []int{1, 2, 3, 4, 4, 7, 7, 12, 12, 13}
	count = countUniqueValues(xi)
	fmt.Println(count) // 7

	xi = []int{}
	count = countUniqueValues(xi)
	fmt.Println(count) // 0

	xi = []int{-2, -1, -1, 0, 1}
	count = countUniqueValues(xi)
	fmt.Println(count) // 4
}

// countUniqueValues accepts a sorted array and counts the number of
// unique values in the array. There can be negative numbers in the array.

func countUniqueValues(xi []int) int {
	if len(xi) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(xi); j++ {
		if xi[i] != xi[j] {
			// first i items would be unique
			i++
			xi[i] = xi[j]
		}
	}
	return i + 1
}
