package main

import (
	"fmt"
	"log"
)

type DoublyNode struct {
	value interface{}
	next  *DoublyNode
	prev  *DoublyNode
}

type DoublyLinkedList struct {
	head   *DoublyNode
	tail   *DoublyNode
	length int
}

func NewDoublyLinkedList(value interface{}) *DoublyLinkedList {
	newNode := &DoublyNode{
		value: value,
		next:  nil,
		prev:  nil,
	}

	return &DoublyLinkedList{
		head:   newNode,
		tail:   newNode,
		length: 1,
	}
}

func (ll *DoublyLinkedList) append(value interface{}) {
	newNode := &DoublyNode{
		value: value,
		next:  nil,
		prev:  nil,
	}

	newNode.prev = ll.tail
	ll.tail.next = newNode
	ll.tail = newNode

	ll.length++
}

func (ll *DoublyLinkedList) prepend(value interface{}) {
	newNode := &DoublyNode{
		value: value,
		next:  nil,
		prev:  nil,
	}

	newNode.next = ll.head
	ll.head.prev = newNode
	ll.head = newNode

	ll.length++
}

func (ll *DoublyLinkedList) insert(index int, value interface{}) {
	if index >= ll.length {
		ll.append(value)
		return
	}

	newNode := &DoublyNode{
		value: value,
		next:  nil,
		prev:  nil,
	}

	// since the index is starting from 0
	leader := ll.traversalToIndex(index - 1)
	follower := leader.next

	leader.next = newNode
	newNode.prev = leader
	newNode.next = follower
	follower.prev = newNode

	ll.length++
}

func (ll *DoublyLinkedList) delete(index int) {
	if index < 0 || index >= ll.length {
		log.Printf("index out of bound: %d", index)
		return
	}

	if index == 0 {
		ll.head = ll.head.next
		ll.head.prev = nil
		if ll.length == 1 {
			ll.tail = nil
		}
	} else {
		leader := ll.traversalToIndex(index - 1)
		unwantedNode := leader.next
		leader.next = unwantedNode.next

		if index == ll.length-1 {
			ll.tail = leader
		}

		if unwantedNode.next != nil {
			unwantedNode.next.prev = leader
		}
	}

	ll.length--
}

func (ll *DoublyLinkedList) traversalToIndex(index int) *DoublyNode {
	currentIndex := 0
	currentNode := ll.head

	for currentIndex != index {
		currentNode = currentNode.next
		currentIndex++
	}

	return currentNode
}

func (ll *DoublyLinkedList) display() {
	current := ll.head
	for current != nil {
		fmt.Printf("Address: %p, Value: %v, Prev: %p, Next: %p -> \n", current, current.value, current.prev, current.next)
		current = current.next
	}
}
