package main

import "fmt"

func main() {
	s1 := "reverse"
	s2 := reverse(s1)
	fmt.Println("result:", s2)
}

// reverse accepts a string and returns a new string in reverse
func reverse(s string) string {
	r := []rune(s)
	r = helper(r)
	return string(r)
}

func helper(xr []rune) []rune {
	if len(xr) <= 1 {
		return xr
	}
	return append(helper(xr[1:]), xr[0])
}
