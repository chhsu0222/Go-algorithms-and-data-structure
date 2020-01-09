package main

import (
	"errors"
	"fmt"
	"math"
)

type maxBinaryHeap struct {
	values []int
}

func (h *maxBinaryHeap) insert(v int) {
	h.values = append(h.values, v)
	h.bubbleUp()
}

// Swap the 1st with the last value in values prop
// Pop from the values prop so we can return the value at the end
// Calls sinkDown
func (h *maxBinaryHeap) extractMax() (int, error) {
	length := len(h.values)
	if length == 0 {
		return 0, errors.New("Heap is empty")
	}

	max := h.values[0]
	if length == 1 {
		h.values = []int{}
		return max, nil
	}
	end := h.values[length-1]
	h.values[0] = end
	h.values = h.values[:length-1]
	h.sinkDown()
	return max, nil
}

// Create a variable call index which is the length of the values prop - 1
// Create a variable call parentIndex which is the floor of (index - 1) / 2
// Keep looping as long as the value at the parentIndex < the value at index
//     Swap the value at parentIndex with the value at index
//     Set the index to be parentIndex
func (h *maxBinaryHeap) bubbleUp() {
	idx := len(h.values) - 1
	child := h.values[idx]
	for idx > 0 {
		parentIdx := int(math.Floor(float64((idx - 1) / 2)))
		parent := h.values[parentIdx]
		if parent >= child {
			break
		}
		h.values[idx] = parent
		h.values[parentIdx] = child
		idx = parentIdx
	}
}

// Parent index start at 0 (the root)
// Find the index of the left child: 2* indx + 1 (make sure it's not out of bounds)
// Find the index of the right child: 2* indx + 2 (make sure it's not out of bounds)
// If the left child > the value, the index of the left index is the candidate
// If the right child > the value, check if it > the left child. If yes, update the candidate
// Swap the value at index with the value at the candidate
// Update the index to the candidate
// Keep looping until the child < value
func (h *maxBinaryHeap) sinkDown() {
	value := h.values[0]
	idx := 0
	length := len(h.values)
	leftIdx, rightIdx := -1, -1
	left, right := 0, 0

	for {
		candidate := -1 // reset candidate for every loop
		leftIdx = 2*idx + 1
		rightIdx = 2*idx + 2

		if leftIdx < length {
			left = h.values[leftIdx]
			if left > value {
				candidate = leftIdx
			}
		}
		if rightIdx < length {
			right = h.values[rightIdx]
			if (candidate == -1 && right > value) || (candidate != -1 && right > left) {
				candidate = rightIdx
			}
		}
		if candidate == -1 {
			break
		}
		h.values[idx] = h.values[candidate]
		h.values[candidate] = value
		idx = candidate
	}
}

func main() {
	h := &maxBinaryHeap{
		values: []int{41, 39, 33, 18, 27, 12},
	}

	h.insert(55)
	fmt.Println(h.values)
	max, err := h.extractMax()
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("max:", max)
	}
	max, err = h.extractMax()
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("max:", max)
	}
	max, err = h.extractMax()
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("max:", max)
	}
	max, err = h.extractMax()
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("max:", max)
	}
	max, err = h.extractMax()
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("max:", max)
	}
	max, err = h.extractMax()
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("max:", max)
	}
	max, err = h.extractMax()
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("max:", max)
	}
	max, err = h.extractMax()
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("max:", max)
	}
}
