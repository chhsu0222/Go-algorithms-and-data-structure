package main

import (
	"errors"
	"fmt"
	"math"
)

// Lower number means higher priority
type node struct {
	value    string
	priority int
}

type priorityQueue struct {
	values []*node
}

// enqueue accepts a value and priority
// makes a new node and puts it in the right spot
// based off of its priority
func (q *priorityQueue) enqueue(v string, p int) {
	n := &node{
		value:    v,
		priority: p,
	}
	q.values = append(q.values, n)
	q.bubbleUp()
}

// dequeue removes root element, returns it and rearanges heap using priority
// Swap the 1st with the last node in values prop
// Pop from the values prop so we can return the node at the end
// Calls sinkDown
func (q *priorityQueue) dequeue() (*node, error) {
	length := len(q.values)
	if length == 0 {
		return nil, errors.New("PriorityQueue is empty")
	}

	min := q.values[0]
	if length == 1 {
		q.values = []*node{}
		return min, nil
	}
	end := q.values[length-1]
	q.values[0] = end
	q.values = q.values[:length-1]
	q.sinkDown()
	return min, nil
}

// Create a variable call index which is the length of the values prop - 1
// Create a variable call parentIndex which is the floor of (index - 1) / 2
// Keep looping as long as the priority at the parentIndex > the value at index
//     Swap the node at parentIndex with the node at index
//     Set the index to be parentIndex
func (q *priorityQueue) bubbleUp() {
	idx := len(q.values) - 1
	child := q.values[idx]
	for idx > 0 {
		parentIdx := int(math.Floor(float64((idx - 1) / 2)))
		parent := q.values[parentIdx]
		if parent.priority <= child.priority {
			break
		}
		q.values[idx] = parent
		q.values[parentIdx] = child
		idx = parentIdx
	}
}

// Parent index start at 0 (the root)
// Find the index of the left child: 2* indx + 1 (make sure it's not out of bounds)
// Find the index of the right child: 2* indx + 2 (make sure it's not out of bounds)
// If the left child's priority < the node's priority, the index of the left index is the candidate
// If the right child's priority < the node's priority, check if it < the left child's priority. If yes, update the candidate
// Swap the node at index with the node at the candidate
// Update the index to the candidate
// Keep looping until the child's priority > node's priority
func (q *priorityQueue) sinkDown() {
	n := q.values[0]
	idx := 0
	length := len(q.values)
	leftIdx, rightIdx := -1, -1
	left, right := &node{}, &node{}

	for {
		candidate := -1 // reset candidate for every loop
		leftIdx = 2*idx + 1
		rightIdx = 2*idx + 2

		if leftIdx < length {
			left = q.values[leftIdx]
			if left.priority < n.priority {
				candidate = leftIdx
			}
		}
		if rightIdx < length {
			right = q.values[rightIdx]
			if (candidate == -1 && right.priority < n.priority) || (candidate != -1 && right.priority < left.priority) {
				candidate = rightIdx
			}
		}
		if candidate == -1 {
			break
		}
		q.values[idx] = q.values[candidate]
		q.values[candidate] = n
		idx = candidate
	}
}

func main() {
	q := &priorityQueue{
		values: []*node{},
	}

	q.enqueue("common cold", 5)
	q.enqueue("gunshot wound", 1)
	q.enqueue("high fever", 4)
	q.enqueue("broken arm", 2)
	q.enqueue("glass in foot", 3)

	fmt.Println(q.values)
	min, err := q.dequeue()
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("min:", min.value, min.priority)
	}
	min, err = q.dequeue()
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("min:", min.value, min.priority)
	}
	min, err = q.dequeue()
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("min:", min.value, min.priority)
	}
	min, err = q.dequeue()
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("min:", min.value, min.priority)
	}
	min, err = q.dequeue()
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("min:", min.value, min.priority)
	}
	min, err = q.dequeue()
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("min:", min.value, min.priority)
	}
}
