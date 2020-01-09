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
	values := breadthFirstSearch(bts)
	fmt.Println("values:", values)
}

// Create a queue as a to-do list and a slice to store the values
// of nodes visited.
// Place the root node in the queue
// Loop as long as there is anything in the queue
//    a. Dequeue a node from the queue and push the value
//       of the node into the slice
//    b. If there is a left prop on the node dequeued - add it
//       to the queue
//    c. If there is a right prop on the node dequeued - add it
//       to the queue
// Return the slice that store the value

func breadthFirstSearch(t *binarySearchTree) []int {
	values := []int{}
	if t.root == nil {
		return values
	}

	queue := []*node{t.root}
	n := &node{}
	for len(queue) != 0 {
		n = queue[0]
		queue = queue[1:]
		values = append(values, n.value)
		if n.left != nil {
			queue = append(queue, n.left)
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
	}
	return values
}
