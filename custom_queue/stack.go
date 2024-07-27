package main

// Define the Stack structure
type Stack struct {
	elements []int
}

// Push adds an element to the stack
func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
		panic("Pop from empty stack")
	}
	value := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return value
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}
