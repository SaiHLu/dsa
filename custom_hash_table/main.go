package main

import (
	"fmt"
	"log"
)

// visual => https://www.cs.usfca.edu/~galles/visualization/OpenHash.html

type HashTable struct {
	data [][][2]interface{}
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		data: make([][][2]interface{}, size),
	}
}

func (h *HashTable) _hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash + int(key[i])*i) % len(h.data)
	}
	return hash
}

func (h *HashTable) Set(key string, value interface{}) {
	address := h._hash(key)
	if h.data[address] == nil {
		h.data[address] = make([][2]interface{}, 0)
	}

	h.data[address] = append(h.data[address], [2]interface{}{
		key,
		value,
	})
}

func (h *HashTable) Get(key string) interface{} {
	address := h._hash(key)
	log.Printf("key %s, address %d", key, address)

	if h.data[address] != nil {
		for _, val := range h.data[address] {
			if val[0] == key {
				return val
			}
		}
	}

	return nil
}

// When having collision, this Keys() method won't work correctly
// According to the fixed index(0) for this example
// Time complexity(collision) - 0(n^2), Time complexity(no collision) - 0(n), Space complexity - O(1)
func (h *HashTable) Keys() []interface{} {
	keys := make([]interface{}, 0)

	for idx := range h.data {
		if len(h.data[idx]) <= 0 {
			continue
		}

		// considering collision case
		for innerIdx := range h.data[idx] {
			if len(h.data[idx][innerIdx]) > 0 {
				keys = append(keys, h.data[idx][innerIdx][0])
			}
		}
		// without considering collision case
		// keys = append(keys, h.data[idx][0][0])
	}

	return keys
}

// Play with the `size` of hashTable for collision case.
func main() {
	myHashTable := NewHashTable(2)
	myHashTable.Set("grapes", 10000)
	myHashTable.Set("apples", 9)
	myHashTable.Set("pineapples", 10)
	myHashTable.Set("snakes", 10)
	fmt.Println(myHashTable.Get("grapes"))     // Output: 10000
	fmt.Println(myHashTable.Get("apples"))     // Output: 9
	fmt.Println(myHashTable.Get("pineapples")) // Output: 10
	fmt.Println(myHashTable.Get("snakes"))     // Output: 10
	fmt.Println("keys: ", myHashTable.Keys())  // Output: [grapes apples pineapples snakes]
}
