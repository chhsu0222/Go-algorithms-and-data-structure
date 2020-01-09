package main

import "fmt"

func main() {
	// xi1 := []int{1, 2, 3}
	// xi2 := []int{9, 1, 4}
	xi3 := []int{3, 1, 5}
	xi4 := []int{9, 1, 21}
	fmt.Println(same(xi3, xi4))
}

// Target:
// Checking if the value in slice1 has the corresponding square value in slice2.
// The number of values should match.
// e.g.
// slice1 = [1, 2, 1] & slice2 = [4, 1, 1] --> match.
// slice1 = [1, 1, 3] & slice2 = [9, 1] --> not match.

func same(xi1, xi2 []int) bool {
	if len(xi1) != len(xi2) {
		return false
	}

	for _, v := range xi1 {
		i := -1
		for j, v2 := range xi2 {
			if v2 == v*v {
				i = j
				break
			}
		}
		if i == -1 {
			return false
		}
		xi2 = append(xi2[:i], xi2[i:]...)
	}
	return true
}
