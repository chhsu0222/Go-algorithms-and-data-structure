package main

import "fmt"

func main() {
	bts := &binarySearchTree{}
	bts.insert(10)
	bts.insert(15)
	bts.insert(6)
	bts.insert(3)
	bts.insert(20)
	bts.insert(8)
	values := depthFirstSearch(bts, "inOrder")
	fmt.Println("values:", values)
}

// Create  a slice to store the values of nodes visited.
// Store the root of BST in a variable called current
// Write a helper function which accepts a node
func depthFirstSearch(t *binarySearchTree, order string) []int {
	values := []int{}
	if t.root == nil {
		return values
	}

	switch order {
	case "preOrder":
		preOrder(t.root, &values)
	case "postOrder":
		postOrder(t.root, &values)
	case "inOrder":
		inOrder(t.root, &values)
	default:
		preOrder(t.root, &values)
	}
	return values
}

// a. If the node has a left prop, call the helper function
//    with the left prop on the node
// b. If the node has a right prop, call the helper function
//    with the right prop on the node
// c. Push the value of the node to the slice
func preOrder(n *node, values *[]int) {
	*values = append(*values, n.value)
	if n.left != nil {
		preOrder(n.left, values)
	}
	if n.right != nil {
		preOrder(n.right, values)
	}
}

// a. If the node has a left prop, call the helper function
//    with the left prop on the node
// b. If the node has a right prop, call the helper function
//    with the right prop on the node
// c. Push the value of the node to the slice
func postOrder(n *node, values *[]int) {
	if n.left != nil {
		postOrder(n.left, values)
	}
	if n.right != nil {
		postOrder(n.right, values)
	}
	*values = append(*values, n.value)
}

// a. If the node has a left prop, call the helper function
//    with the left prop on the node
// b. Push the value of the node to the slice
// c. If the node has a right prop, call the helper function
//    with the right prop on the node
func inOrder(n *node, values *[]int) {
	if n.left != nil {
		inOrder(n.left, values)
	}

	*values = append(*values, n.value)

	if n.right != nil {
		inOrder(n.right, values)
	}
}
