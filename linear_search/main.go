package main

import (
	"log"
)

func main() {
	log.Println("linear search")
}

func linearSearch(haystack []int, needle int) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}

	return false
}
