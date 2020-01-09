package main

import (
	"errors"
	"fmt"
)

type node struct {
	value string
	next  *node
	pre   *node
}

type doublyLinkedList struct {
	head   *node
	tail   *node
	length int
}

// push accepts a string and creates a new node using the value
// passed to the function.
// If there is no head, set the head and tail to be the newly created node.
// Otherwise set the next prop on the tail to be the new node
// Set the pre prop on the new node to be the tail
// Set the tail property on the list to be the newly created node.
// Increment the length by 1.
func (d *doublyLinkedList) push(v string) {
	n := &node{
		value: v,
		next:  nil,
		pre:   nil,
	}

	if d.head == nil {
		d.head = n
		d.tail = n
	} else {
		d.tail.next = n
		n.pre = d.tail
		d.tail = n
	}
	d.length++
}

// pop remove the last node of the list.
// If there are no nodes in the list, return an error message.
// Store the current tail in a variable to return later
// If the length is 1, set the head and tail to be nill
// Update the tail to be the previous node
// Set the new tail's next to be nil
// Set the removed node's pre to be nil
// Decrement the length of the list by 1.
// Return the value of the node removed
func (d *doublyLinkedList) pop() (string, error) {
	if d.head == nil {
		return "", errors.New("List is empty")
	}
	n := d.tail
	if d.length == 1 {
		d.head = nil
		d.tail = nil
	} else {
		d.tail = d.tail.pre
		d.tail.next = nil
		n.pre = nil
	}
	d.length--
	return n.value, nil
}

// shift removes a node from the beginning of the list
// If there are no nodes, return an error message.
// Store the current head prop in a variable
// If the length is 1, set the head and tail to be nill
// Update the head to be the current head's next prop
// Set the pre prop of the head to be nil
// Set the next prop of the removed node to be nil
// Decrement the length by 1
// Return the value of the node removed
func (d *doublyLinkedList) shift() (string, error) {
	if d.head == nil {
		return "", errors.New("List is empty")
	}
	n := d.head
	if d.length == 1 {
		d.head = nil
		d.tail = nil
	} else {
		d.head = d.head.next
		d.head.pre = nil
		n.next = nil
	}
	d.length--
	return n.value, nil
}

// unshift accepts a value and adds a new node to the beginning of the list.
// Create new node using the value passed to the function
// If there is no head property on the list, set the head and tail to be the newly created node.
// Otherwise, set the pre prop of the head to be the new node
// set the newly created node's next prop to be the current head prop on the list
// Set the head prop on the list to be that newly created node
// Increment the length of the list by 1
func (d *doublyLinkedList) unshift(v string) {
	n := &node{
		value: v,
		next:  nil,
		pre:   nil,
	}
	if d.head == nil {
		d.head = n
		d.tail = n
	} else {
		d.head.pre = n
		n.next = d.head
		d.head = n
	}
	d.length++
}

// get retrives a node by it's position in the list
// If the index is less than 0 or >= length, return an error message.
// If the index is <= (length / 2)
//     Loop through the list starting from the head to the middle
//     Return the node once it is found
// else
//     Loop through the list starting from the tail to the middle
//     Return the node once it is found
func (d *doublyLinkedList) get(i int) (*node, error) {
	if i < 0 || i >= d.length {
		return nil, errors.New("Invalid index")
	}
	current := d.head
	if i <= d.length/2 {
		for count := 0; count < i; count++ {
			current = current.next
		}
	} else {
		current = d.tail
		for count := d.length - 1; count > i; count-- {
			current = current.pre
		}
	}
	return current, nil
}

// set sets the value of a node based on it's position in the list
// It accepts a value and an index.
// Using the get method to get the node.
// If the get method returns a valid node, Set the value of the node and return true
// Otherwise, return false
func (d *doublyLinkedList) set(i int, v string) bool {
	node, err := d.get(i)
	if err != nil {
		return false
	}
	node.value = v
	return true
}

// insert adds a node to the list at a specific position.
// If the index is < 0 or > the length, return false
// If index == the length, using push
// If index == 0, using unshift
// Otherwise, using the get method, access the node at the index - 1
// Create a new node.
// Set the next prop on the new node to be the previous' next
// Set the pre of the new node to the previous node
// Set the next prop on previous node to be the new node
// Set the pre of the next node to the new node
// Increment the length and return true
func (d *doublyLinkedList) insert(i int, v string) bool {
	if i < 0 || i > d.length {
		return false
	}
	if i == 0 {
		d.unshift(v)
		return true
	}
	if i == d.length {
		d.push(v)
		return true
	}

	previous, _ := d.get(i - 1)
	n := &node{
		value: v,
		next:  previous.next,
		pre:   previous,
	}
	previous.next = n
	n.next.pre = n
	d.length++
	return true
}

func (d *doublyLinkedList) print() {
	fmt.Println("Print!!!")
	current := d.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func (d *doublyLinkedList) printBackward() {
	fmt.Println("Print Backward!!!")
	current := d.tail
	for current != nil {
		fmt.Println(current.value)
		current = current.pre
	}
}

// Remove removes a node from the list at a specific position
// If the index is < 0 or >= the length, return an error message
// If the index is length - 1, using pop
// If the index is 0, using shift
// Otherwise, using the get method, access the node at index
// Set the next prop of the previous node to be the next of the removed node
// Set the pre prop of the next node to be the previous node
// Set the pre & next prop of the removed node to be nil
// Decrement the length and return the value of the node removed
func (d *doublyLinkedList) remove(i int) (string, error) {
	if i < 0 || i >= d.length {
		return "", errors.New("Invalid index")
	}
	if i == d.length-1 {
		return d.pop()
	}
	if i == 0 {
		return d.shift()
	}
	removed, _ := d.get(i)
	removed.pre.next = removed.next
	removed.next.pre = removed.pre
	removed.next = nil
	removed.pre = nil
	d.length--
	return removed.value, nil
}

func main() {
	l := &doublyLinkedList{}
	l.push("hello")
	l.push("are")
	l.unshift("you")
	l.unshift("ok")
	l.print()
	value, err := l.remove(5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("removed value", value)
	}
	value, err = l.remove(2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("removed value", value)
	}
	l.printBackward()
}
