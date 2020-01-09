package main

import "fmt"

func main() {
	s1 := "reverse"
	fmt.Println(isPalindrome(s1)) // false
	s1 = "tacocat"
	fmt.Println(isPalindrome(s1)) // true
}

// isPalindrome accepts a string and returns true if it's a palindrome
// (reads the same forward and backward). Otherwise it returns false.
func isPalindrome(s string) bool {
	return helper([]rune(s))
}

func helper(xr []rune) bool {
	if len(xr) <= 1 {
		return true
	}
	if len(xr) == 2 {
		return xr[0] == xr[1]
	}
	return xr[0] == xr[len(xr)-1] && helper(xr[1:len(xr)-1])
}
