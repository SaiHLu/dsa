package main

import (
	"fmt"
	"log"
)

type Node struct {
	value interface{}
	next  *Node
}

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

func (q *Queue) peek() *Node {
	return q.head
}

func (q *Queue) enqueue(value interface{}) {
	newNode := &Node{
		value: value,
		next:  nil,
	}

	if q.length == 0 {
		q.head = newNode
		q.tail = newNode

	} else {
		q.tail.next = newNode
		q.tail = newNode
	}

	q.length++

}

func (q *Queue) dequeue() {
	if q.head == nil {
		return
	}

	if q.head == q.tail {
		q.tail = nil
	}

	q.head = q.head.next

	q.length--
}

func (q *Queue) display() {
	log.Println("=== HEAD ===")
	currentHead := q.head
	for currentHead != nil {
		fmt.Printf("Address: %p, Value: %v, Next: %p\n", currentHead, currentHead.value, currentHead.next)
		currentHead = currentHead.next
	}

	log.Println("=== TAIL ===")
	currentTail := q.tail
	if currentTail != nil {
		fmt.Printf("Address: %p, Value: %v, Next: %p\n", currentTail, currentTail.value, currentTail.next)
	}

	fmt.Printf("Total Length: %d\n", q.length)
}
