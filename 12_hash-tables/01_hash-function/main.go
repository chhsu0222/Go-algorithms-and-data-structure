package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(basicHash("pink", 20)) // 10
	fmt.Println(basicHash("Pink", 20)) // -2
	fmt.Println(basicHash("pink", 20)) // 10
}

// basicHash accepts a key, which is a string
// and the length of the underling array.
// It will return the index that the corresponding
// value of the key would be stored in the array
func basicHash(key string, arrLen int) int {
	total := 0
	const weirdPrime = 31
	// For the const calculation time, we only consider at most 100
	// characters of the key
	maxLen := int(math.Min(float64(len(key)), 100.0))
	for i := 0; i < maxLen; i++ {
		value := int(key[i]) - 96 // UTF-8
		// Using the prime numbe to increase randomness
		total = (total*weirdPrime + value) % arrLen
	}
	return total
}
