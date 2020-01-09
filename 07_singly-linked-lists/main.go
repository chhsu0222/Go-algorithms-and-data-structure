package main

import (
	"errors"
	"fmt"
)

type node struct {
	value string
	next  *node
}

type singlyLinkedList struct {
	head   *node
	tail   *node
	length int
}

// push accepts a string and creates a new node using the value
// passed to the function. If there is no head, set the head
// and tail to be the newly created node. Otherwise set the next
// prop on the tail to be the new node and set the tail property
// on the list to be the newly created node. Increment the length by 1.
func (s *singlyLinkedList) push(v string) {
	n := &node{
		value: v,
		next:  nil,
	}
	if s.head == nil {
		s.head = n
		s.tail = n
	} else {
		s.tail.next = n
		s.tail = n
	}
	s.length++
}

// pop remove the last node of the list. If there are no nodes
// in the list, return an error message.
// Loop through the list until we reach the tail
// Set the next property of the 2nd to last node to be nil
// Set the tail to be the 2nd to last node
// Decrement the length of the list by 1.
// Return the value of the node removed
func (s *singlyLinkedList) pop() (string, error) {
	if s.head == nil {
		return "", errors.New("List is empty")
	}
	current := s.head
	newTail := current
	for current.next != nil {
		newTail = current
		current = current.next
	}
	newTail.next = nil
	s.tail = newTail
	s.length--
	if s.length == 0 {
		s.head = nil
		s.tail = nil
	}
	return current.value, nil
}

// shift removes a node from the beginning of the list
// If there are no nodes, return an error message.
// Store the current head prop in a variable
// Set the head prop to be the current head's next prop
// Decrement the length by 1
// Return the value of the node removed
func (s *singlyLinkedList) shift() (string, error) {
	if s.head == nil {
		return "", errors.New("List is empty")
	}
	value := s.head.value
	s.head = s.head.next
	s.length--
	if s.length == 0 {
		s.tail = nil
	}
	return value, nil
}

// unshift accepts a value and adds a new node to the beginning of the list.
// Create new node using the value passed to the function
// If there is no head property on the list, set the head and tail to be the newly created node.
// Otherwise set the newly created node's next prop to be the current head prop on the list
// Set the head prop on the list to be that newly created node
// Increment the length of the list by 1
func (s *singlyLinkedList) unshift(v string) {
	n := &node{
		value: v,
		next:  nil,
	}
	if s.head == nil {
		s.head = n
		s.tail = n
	} else {
		n.next = s.head
		s.head = n
	}
	s.length++
}

// get retrives a node by it's position in the list
// It accepts an index. If the index is less than 0 or greater
// than or equal to the length of the list, return an error message.
// Loop through the list until you reach the index and return
// the node at that specific index
func (s *singlyLinkedList) get(i int) (*node, error) {
	if i < 0 || i >= s.length {
		return nil, errors.New("Invalid index")
	}
	current := s.head
	for count := 0; count < i; count++ {
		current = current.next
	}
	return current, nil
}

// set sets the value of a node based on it's position in the list
// It accepts a value and an index. Using the get method to get the node.
// Loop through the list until you reach the index and set the value of
// the node at that specific index
func (s *singlyLinkedList) set(i int, v string) bool {
	node, err := s.get(i)
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
// Set the next prop on that node to be the new node
// Set the next prop on the new node to be the previous next
// Increment the length and return true
func (s *singlyLinkedList) insert(i int, v string) bool {
	if i < 0 || i > s.length {
		return false
	}
	if i == 0 {
		s.unshift(v)
		return true
	}
	if i == s.length {
		s.push(v)
		return true
	}

	pre, _ := s.get(i - 1)
	n := &node{
		value: v,
		next:  pre.next,
	}
	pre.next = n
	s.length++
	return true
}

// Remove removes a node from the list at a specific position
// If the index is < 0 or >= the length, return an error message
// If the index is length - 1, using pop
// If the index is 0, using shift
// Otherwise, using the get method, access the node at index - 1
// Set the next prop on that node to be the next of the next node
// Decrement the length and return the value of the node removed
func (s *singlyLinkedList) remove(i int) (string, error) {
	if i < 0 || i >= s.length {
		return "", errors.New("Invalid index")
	}
	if i == s.length-1 {
		return s.pop()
	}
	if i == 0 {
		return s.shift()
	}
	pre, _ := s.get(i - 1)
	removed := pre.next
	pre.next = removed.next
	removed.next = nil
	s.length--
	return removed.value, nil
}

// reverse the list in place
func (s *singlyLinkedList) reverse() {
	if s.length <= 1 {
		return
	}
	// swap head and tail
	s.head, s.tail = s.tail, s.head
	pre := s.tail
	current := pre.next
	next := current
	for current != nil {
		next = current.next // move forward
		current.next = pre  // point back
		pre = current       // move forward
		current = next      // move forward
	}
	s.tail.next = nil
}

func (s *singlyLinkedList) print() {
	fmt.Println("Print!!!")
	current := s.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func main() {
	list := &singlyLinkedList{}
	list.push("hello")
	list.push("how")
	list.push("are")

	list.print()
	list.reverse()
	list.print()
}
