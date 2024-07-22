package main

import "log"

// Given
var (
	arr1 = []int{2, 5, 1, 2, 3, 5, 1, 2, 4}
	arr2 = []int{2, 5, 5, 2, 3, 5, 1, 2}
	arr3 = []int{2, 3, 4, 5}
)

func main() {
	result := findFirstRepeatedNumber(arr2, 3)
	log.Println(result)
}

// params - numbers []int
// params - times => repeated times
func findFirstRepeatedNumber(numbers []int, times int) int {
	result := make(map[int]int)

	if times <= 0 {
		return -1
	}

	for _, val := range numbers {
		result[val]++

		if result[val] == times {
			return val
		}
	}

	return -1
}
