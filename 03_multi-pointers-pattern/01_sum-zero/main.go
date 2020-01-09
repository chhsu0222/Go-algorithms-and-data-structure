package main

import "fmt"

func main() {
	xi := []int{-5, -4, -2, -1, 0, 1, 2, 3, 10}
	pair := sumZero(xi)
	fmt.Println(pair) // [-2, 2]

	xi = []int{0, 1, 2, 3, 10}
	pair = sumZero(xi)
	fmt.Println(pair) // []
}

// sumZero accepts a sorted array and finds the first pair
// where the sum is 0. Return an array that includes both
// values that sum to zero or an empty array.
func sumZero(xi []int) []int {
	left := 0
	right := len(xi) - 1
	pair := make([]int, 0, 2)

	sum := -1
	for left < right {
		sum = xi[left] + xi[right]
		if sum == 0 {
			pair = append(pair, xi[left], xi[right])
			return pair
		} else if sum > 0 {
			right--
		} else {
			left++
		}
	}
	return pair
}
