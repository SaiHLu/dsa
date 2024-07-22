package main

import (
	"log"
	"math"
)

func main() {
	log.Println("Binary Search ", math.Sqrt(5))
}

func binarySearch(haystack []int, needle int) bool {
	var (
		lowerBound float64 = 0
		upperBound         = float64(len(haystack))
	)

	for lowerBound < upperBound {
		middleIndex := math.Floor((lowerBound + upperBound) / 2)
		// middleIndex := math.Floor(lowerBound + (upperBound-lowerBound)/2)
		middleValue := haystack[int(middleIndex)]

		if middleValue == needle {
			return true
		} else if middleValue < needle {
			lowerBound = middleIndex + 1
		} else {
			upperBound = middleIndex
		}
	}

	return false
}
