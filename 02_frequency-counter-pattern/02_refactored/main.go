package main

import "fmt"

func main() {
	xi1 := []int{1, 2, 3}
	xi2 := []int{9, 1, 4}
	fmt.Println(same(xi1, xi2))
	// xi3 := []int{3, 1, 5}
	// xi4 := []int{9, 1, 21}
	// fmt.Println(same(xi3, xi4))
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

	m1 := make(map[int]int)
	m2 := make(map[int]int)

	for _, v := range xi1 {
		m1[v]++
	}
	for _, v := range xi2 {
		m2[v]++
	}

	for k, v := range m1 {
		if m2[k*k] != v {
			return false
		}
	}
	return true
}
