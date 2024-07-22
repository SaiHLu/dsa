package main

import "log"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	expectedResult := 9

	log.Println("match: ", hasParisWithSum2(arr, expectedResult))
}

// Bad
func hasPairsWithSum(arr []int, result int) bool {
	for _, val := range arr {
		for _, val2 := range arr {
			if val+val2 == result {
				return true
			}
		}
	}
	return false
}

// Good
func hasParisWithSum2(arr []int, result int) bool {
	tmpMap := make(map[int]bool)

	for _, val := range arr {
		tmpMap[result-val] = true

		if tmpMap[val] {
			return true
		}
	}

	return false
}
