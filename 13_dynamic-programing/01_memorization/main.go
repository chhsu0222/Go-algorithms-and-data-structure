package main

import "fmt"

func main() {
	m := make(map[int]int)
	result := fib(5, m)
	fmt.Println(result)
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// Memorization
// Stroing the result of expensive function calls
// and returning the cashed result when the same
// inputs occur again
func fib(n int, m map[int]int) int {
	if v, ok := m[n]; ok {
		return v
	}
	if n <= 2 {
		return 1
	}
	result := fib(n-1, m) + fib(n-2, m)
	m[n] = result
	return result
}
