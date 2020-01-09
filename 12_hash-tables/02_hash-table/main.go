package main

import (
	"errors"
	"fmt"
	"math"
)

type hashTable struct {
	keyMap []*[][]string
}

func (ht *hashTable) _init(size int) {
	ht.keyMap = make([]*[][]string, size)
	for i := range ht.keyMap {
		ht.keyMap[i] = &[][]string{}
	}
}

func (ht *hashTable) _hash(key string) int {
	total := 0
	const weirdPrime = 31
	maxLen := int(math.Min(float64(len(key)), 100.0))
	for i := 0; i < maxLen; i++ {
		value := int(key[i]) - 96
		total = (total*weirdPrime + value) % len(ht.keyMap)
	}
	return total
}

// set accepts a key and a value
// Hashes the key
// Stores the key-value pair in the hash table array via separate chaining
func (ht *hashTable) set(key, value string) {
	idx := ht._hash(key)
	*ht.keyMap[idx] = append(*ht.keyMap[idx], []string{key, value})
}

// get accepts a key
// Hashes the key
// Retrive the key-value pair in the hash table
func (ht *hashTable) get(key string) (string, error) {
	idx := ht._hash(key)
	for _, v := range *ht.keyMap[idx] {
		if v[0] == key {
			return v[1], nil
		}
	}
	return "", errors.New("undefined")
}

// Loops through the hash table and returns an slice
// of unique keys in the table
func (ht *hashTable) keys() []string {
	keys := []string{}
	for _, bucket := range ht.keyMap {
		for _, pair := range *bucket {
			if !contains(keys, pair[0]) {
				keys = append(keys, pair[0])
			}
		}
	}
	return keys
}

// Loops through the hash table and returns an slice
// of unique values in the table
func (ht *hashTable) values() []string {
	values := []string{}
	for _, bucket := range ht.keyMap {
		for _, pair := range *bucket {
			if !contains(values, pair[1]) {
				values = append(values, pair[1])
			}
		}
	}
	return values
}

func contains(xs []string, s string) bool {
	for _, v := range xs {
		if v == s {
			return true
		}
	}
	return false
}

func main() {
	ht := &hashTable{}
	ht._init(17)
	ht.set("marron", "#800000")
	ht.set("yellow", "FFFF00")
	ht.set("olive", "#808000")
	ht.set("salmon", "#FA8072")
	ht.set("lightcoral", "#F08080")
	ht.set("mediumvioletred", "#C71585")
	ht.set("plum", "#DDA0DD")
	for _, v := range ht.keyMap {
		fmt.Println(*v)
	}

	keys := ht.keys()
	fmt.Println(keys)
	values := ht.values()
	fmt.Println(values)
}
