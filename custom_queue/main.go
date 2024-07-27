package main

import (
	"log"
)

func main() {
	queue := &Queue{}

	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)
	queue.dequeue()
	log.Println(queue.peek())
	queue.dequeue()
	queue.dequeue()
	queue.dequeue()
	queue.dequeue()

	queue.display()
}
