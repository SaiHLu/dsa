package main

import "log"

// LIFO
func main() {
	stack := &StackWithLinkedList{}
	stack.push("google")
	stack.push("udemy")
	stack.push("discord")
	stack.display()
	log.Println("---")
	stack.pop()
	stack.pop()
	stack.pop()
	stack.display()

	// stackWithArr := NewStackWithArray()
	// stackWithArr.push("google")
	// stackWithArr.push("udemy")
	// stackWithArr.push("discord")
	// log.Println("--- Push ---")
	// log.Println(stackWithArr.arr)
	// log.Println("--- Pop ---")
	// stackWithArr.pop()
	// log.Println(stackWithArr.arr)
	// log.Println("--- Peek ---")
	// log.Println(stackWithArr.peek())
	// log.Println(stackWithArr.arr)
}
