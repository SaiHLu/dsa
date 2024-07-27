package main

import (
	"fmt"
	"log"
)

type Node struct {
	value interface{}
	next  *Node
}

// LIFO => last-in-first-out
type StackWithLinkedList struct {
	top    *Node
	bottom *Node
	length int
}

func (s *StackWithLinkedList) push(value interface{}) {
	newNode := &Node{
		value: value,
		next:  nil,
	}

	if s.length == 0 {
		s.top = newNode
		s.bottom = newNode
	}

	if s.length > 0 {
		leader := s.top
		s.top = newNode
		s.top.next = leader
	}

	s.length++
}

func (s *StackWithLinkedList) pop() {
	if s.top == nil {
		return
	}

	if s.top == s.bottom {
		s.bottom = nil
	}

	follower := s.top.next
	s.top = follower
	s.length--
}

func (s *StackWithLinkedList) peek() *Node {
	return s.top
}

func (s *StackWithLinkedList) display() {
	currentTop := s.top

	log.Println("=== Top ===")
	for currentTop != nil {
		fmt.Printf("Address: %p, Value: %v,  Next: %p -> \n", currentTop, currentTop.value, currentTop.next)
		currentTop = currentTop.next
	}

	log.Println("=== Bottom ===")
	currentBottom := s.bottom
	if currentBottom != nil {
		fmt.Printf("Address: %p, Value: %v,  Next: %p -> \n", currentBottom, currentBottom.value, currentBottom.next)
	}
}
