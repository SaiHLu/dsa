package main

import (
	"log"
)

type Array struct {
	length int
	data   map[int]interface{}
}

func (a Array) get(idx int) interface{} {
	return a.data[idx]
}

func (a *Array) push(data interface{}) {
	a.data[a.length] = data
	a.length++
}

func (a *Array) pop() interface{} {
	lastItem := a.data[a.length-1]
	delete(a.data, a.length-1)
	a.length--

	return lastItem
}

func (a *Array) delete(idx int) {
	delete(a.data, idx)
	a.reIndex(idx)
}

func (a *Array) reIndex(idx int) {
	for i := idx; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
	delete(a.data, a.length-1)
	a.length--
}

func main() {
	arr1 := &Array{
		length: 0,
		data:   make(map[int]interface{}),
	}

	arr1.push(1)
	arr1.push(2)
	arr1.push(3)
	arr1.push("four")

	log.Println(arr1)
	arr1.delete(3)
	log.Println(arr1)
}
