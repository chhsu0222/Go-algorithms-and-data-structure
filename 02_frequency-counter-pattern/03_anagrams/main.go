package main

import "fmt"

func main() {
	s1 := ""
	s2 := ""
	fmt.Println(validAnagram(s1, s2))
	s1 = "aaz"
	s2 = "zza"
	fmt.Println(validAnagram(s1, s2))
	s1 = "anagram"
	s2 = "nagaram"
	fmt.Println(validAnagram(s1, s2))
	s1 = "rat"
	s2 = "car"
	fmt.Println(validAnagram(s1, s2))
}

// Target:
// Checking if the 2nd string is an anagram of the first string.
// e.g.
// string1 = "" & string2 = "" --> true
// string1 = "aaz" & string2 = "zza" --> false
// string1 = "anagram" & string2 = "nagaram" --> true
// string1 = "rat" & string2 = "car" --> false
// string1 = "qwerty" & string2 = "qeywrt" --> true

func validAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m := make(map[string]int)

	for _, v := range s1 {
		m[string(v)]++
	}
	for _, v := range s2 {
		if m[string(v)] == 0 {
			// can't find char or char is zero then it's not an anagram
			return false
		}
		m[string(v)]--
	}
	return true
}
