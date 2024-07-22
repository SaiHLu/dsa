package main

import "log"

type Node struct {
	value interface{}
	next  *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedList(value interface{}) *LinkedList {
	newNode := &Node{
		value: value,
		next:  nil,
	}

	return &LinkedList{
		head:   newNode,
		tail:   newNode,
		length: 1,
	}
}

func (ll *LinkedList) append(value interface{}) {
	newNode := &Node{
		value: value,
		next:  nil,
	}

	ll.tail.next = newNode
	ll.tail = newNode

	ll.length++
}

func (ll *LinkedList) prepend(value interface{}) {
	newNode := &Node{
		value: value,
		next:  nil,
	}

	newNode.next = ll.head
	ll.head = newNode

	ll.length++
}

func (ll *LinkedList) insert(index int, value interface{}) {
	if index >= ll.length {
		ll.append(value)
		return
	}

	newNode := &Node{
		value: value,
		next:  nil,
	}

	// since the index is starting from 0
	leader := ll.traversalToIndex(index - 1)
	holdingPointer := leader.next

	leader.next = newNode
	newNode.next = holdingPointer

	ll.length++
}

func (ll *LinkedList) delete(index int) {
	if index < 0 || index >= ll.length {
		log.Printf("index out of bound: %d", index)
		return
	}

	if index == 0 {
		ll.head = ll.head.next
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
	}

	ll.length--
}

func (ll *LinkedList) traversalToIndex(index int) *Node {
	currentIndex := 0
	currentNode := ll.head

	for currentIndex != index {
		currentNode = currentNode.next
		currentIndex++
	}

	return currentNode
}

func (ll *LinkedList) display() {
	current := ll.head
	for current != nil {
		log.Printf("%v", current.value)
		current = current.next
	}
}

func main() {
	ll := NewLinkedList(0)

	ll.append(1)
	ll.append(2)
	ll.append(3)
	ll.append(4)
	ll.append(5)

	ll.delete(5)
	ll.display()
}
