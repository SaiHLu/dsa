package main

type StackWithArray struct {
	arr []interface{}
}

func NewStackWithArray() *StackWithArray {
	return &StackWithArray{
		arr: make([]interface{}, 0),
	}
}

func (s *StackWithArray) push(value interface{}) {
	s.arr = append(s.arr, value)
}

func (s *StackWithArray) pop() {
	if len(s.arr) <= 0 {
		return
	}

	s.arr = s.arr[:len(s.arr)-1]
}

func (s *StackWithArray) peek() interface{} {
	if len(s.arr) <= 0 {
		return "no value"
	}

	return s.arr[len(s.arr)-1]
}
