package main

import "fmt"

func main() {
	result := fib(10)
	fmt.Println(result)
}

// Tabulation (a bottom up approach)
// Storing the result of a previous result in a table (usually an list)
// Usually done using iteration
// Better space complexity
func fib(n int) int {
	if n <= 2 {
		return 1
	}
	fibNums := []int{0, 1, 1}
	for i := 3; i <= n; i++ {
		fibNums = append(fibNums, fibNums[i-1]+fibNums[i-2])
	}
	return fibNums[n]
}
