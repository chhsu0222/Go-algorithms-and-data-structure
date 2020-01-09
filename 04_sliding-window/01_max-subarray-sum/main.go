package main

import "fmt"

func main() {
	xi := []int{1, 2, 5, 2, 8, 1, 5}
	maxSum := maxSubarraySum(xi, 2)
	fmt.Println(maxSum) // 10

	xi = []int{1, 2, 5, 2, 8, 1, 5}
	maxSum = maxSubarraySum(xi, 4)
	fmt.Println(maxSum) // 17

	xi = []int{}
	maxSum = maxSubarraySum(xi, 4)
	fmt.Println(maxSum) // -1

	xi = []int{4, 2, 1, 6, 2}
	maxSum = maxSubarraySum(xi, 4)
	fmt.Println(maxSum) // 13
}

// maxSubarraySum accepts an array and a number n. It will calculate
// the maximum sum of n consecutive elements in the array.

func maxSubarraySum(xi []int, n int) int {
	if len(xi) < n {
		return -1
	}

	maxSum := 0
	for i := 0; i < n; i++ {
		maxSum += xi[i]
	}

	tempSum := maxSum
	for j := n; j < len(xi); j++ {
		tempSum = tempSum - xi[j-n] + xi[j]
		if tempSum > maxSum {
			maxSum = tempSum
		}
	}
	return maxSum
}
