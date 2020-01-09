package main

import (
	"fmt"
)

func main() {
	s := "Your PIN number!!!"
	fmt.Println(strCount(s))
}

func strCount(str string) map[string]int {
	// make a map to return at end
	m := make(map[string]int)
	// loop over string, for each character...
	for _, v := range str {
		// if the char is a number/letter, add one to count
		if isAlphaNumeric(v) {
			// To lower case
			if v > 64 && v < 91 {
				v += 32
			}
			m[string(v)]++
		}
	}
	// return map at end
	return m
}

// Check if char is something else (space, period, ect.)
func isAlphaNumeric(r rune) bool {
	if !(r > 47 && r < 58) && // numeric (0-9)
		!(r > 64 && r < 91) && // upper alpha (A-Z)
		!(r > 96 && r < 123) { // lower alpha (a-z)
		return false
	}
	return true
}
