package main

import (
	"log"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	needle := 11

	result := binarySearch(haystack, needle)

	if !result {
		t.Errorf("Could not find value: %d", needle)
	}

	log.Println("Find value: ", needle)
}
