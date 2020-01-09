package main

import (
	"fmt"
	"math"
)

func main() {
	xi := []int{23, 9852, 12, 5467, 345, 2345}
	// fmt.Println(getDigit(5436, 0))
	// fmt.Println(digitCount(5436))
	// fmt.Println(maxDigits([]int{248, 135, 2, 5555}))
	xi = radixSort(xi)
	fmt.Println(xi)
}

// getDigit gets the ith position of n, which is a base 10 number.
// i starts from 0
func getDigit(n, i int) int {
	return int(math.Floor(math.Abs(float64(n))/math.Pow10(i))) % 10
}

// digitCount counts the digits of n, which is a base 10 number
func digitCount(n int) int {
	if n == 0 {
		return 1
	}
	return int(math.Floor(math.Log10(math.Abs(float64(n))))) + 1
}

// maxDigits accepts a slice of int and returns the number of digits
// in the largest numbers in the slice
func maxDigits(xi []int) int {
	max := 0
	for _, v := range xi {
		max = int(math.Max(float64(max), float64(digitCount(v))))
	}
	return max
}

func radixSort(xi []int) []int {
	// Figure out how many digits the largest numbe has
	max := maxDigits(xi)
	for k := 0; k < max; k++ {
		// Create buckets for each digit (0 - 9)
		digitBuckets := make([][]int, 10)
		for i := range digitBuckets {
			digitBuckets[i] = make([]int, 0)
		}
		// Place each number in the corresponding bucket based on its kth digit
		for _, v := range xi {
			digit := getDigit(v, k)
			digitBuckets[digit] = append(digitBuckets[digit], v)
		}
		fmt.Println("digitBuckets:", digitBuckets)
		// Replace our existing array with values in our buckets, starting with 0 and going up to 9
		xi = make([]int, 0, len(xi))
		for _, bucket := range digitBuckets {
			xi = append(xi, bucket...)
		}
		fmt.Println("xi:", xi)
	}
	return xi
}
