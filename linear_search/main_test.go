package main

import (
	"log"
	"os"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5}
	needle := 30

	exists := linearSearch(haystack, needle)
	if !exists {
		log.Println("Not Exists")
		os.Exit(1)
	}

	log.Println("exists")
}
