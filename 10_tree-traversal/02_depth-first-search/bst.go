package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

type binarySearchTree struct {
	root *node
}

// insert accepts a value
// Create a new node
// Start at the root
// Check if there is a root, if not, root now become that new node
// If there is a root, check if the value of the new node is > or < the value of the root
// If > the value of the root
//    Check if there is a node to the right
//        If there is, move to that node and repeat these steps
//        If there is not, add that node as the right prop
// If < the value of the root
//    Check if there is a node to the left
//        If there is, move to that node and repeat these steps
//        If there is not, add that node as the left prop
func (bts *binarySearchTree) insert(v int) {
	n := &node{
		value: v,
		left:  nil,
		right: nil,
	}
	if bts.root == nil {
		bts.root = n
		return
	}

	current := bts.root
	for {
		if v > current.value {
			if current.right == nil {
				current.right = n
				return
			}
			current = current.right
		} else if v < current.value {
			if current.left == nil {
				current.left = n
				return
			}
			current = current.left
		} else {
			fmt.Println("The value already existed in the BTS.")
			return
		}
	}
}

// find accepts a value
// Start at the root
// Check if there is a root, if not, we're done searching
// If there is a root, check if the value == value of the root
// If > the value of the root
//    Check if there is a node to the right
//        If there is, move to that node and repeat these steps
//        If there is not, we're done searching
// If < the value of the root
//    Check if there is a node to the left
//        If there is, move to that node and repeat these steps
//        If there is not, we're done searching
func (bts *binarySearchTree) find(v int) bool {
	if bts.root == nil {
		return false
	}

	current := bts.root
	for {
		if v == current.value {
			return true
		} else if v > current.value {
			if current.right == nil {
				return false
			}
			current = current.right
		} else if v < current.value {
			if current.left == nil {
				return false
			}
			current = current.left
		}
	}
}
