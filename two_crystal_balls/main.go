// NOTE: FAILED ALGO
package main

import (
	"log"
	"math"
)

// The "two crystal balls" problem involves finding the highest floor from which you can drop a crystal ball without it breaking, using as few drops as possible.
func main() {
	log.Println("Two Crystal Balls")
}

// O(sqrt(n))
func breakingPoints(breaks []bool) int64 {
	// get jumpPoint as sqrt of breaks length
	jumpAmount := math.Floor(math.Sqrt(float64(len(breaks))))
	var breakIndex float64

	for breakIndex = jumpAmount; breakIndex < float64(len(breaks)); breakIndex += jumpAmount {
		if breaks[int(breakIndex)] {
			break
		}
	}

	breakIndex -= jumpAmount

	for j := 0; j < int(jumpAmount) && breakIndex < float64(len(breaks)); j++ {
		if breaks[int64(breakIndex)] {
			return int64(breakIndex)
		}
		breakIndex++
	}

	return -1
}
