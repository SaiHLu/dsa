package main

import "log"

func main() {
	ll := NewLinkedList(0)
	ll.append(1)
	log.Println("original...")
	ll.display()
	log.Println("reverse...")
	ll.reverse()
	ll.display()
	log.Println("=======")

	dll := NewDoublyLinkedList(0)
	dll.append(1)
	dll.append(2)
	dll.prepend(3)
	dll.insert(1, "Insert")
	dll.delete(1)
	dll.display()

}
