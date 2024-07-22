package main

import "log"

func main() {
	arr1 := []interface{}{"1", "2", "3", "1"}
	arr2 := []interface{}{"4", "5", "1"}
	// arr1 := []interface{}{1, 2, 3}
	// arr2 := []interface{}{4, 5, 1}

	log.Println(containCommon(arr1, arr2))
}

func containCommon(arr1, arr2 []interface{}) bool {
	var resultMap = make(map[interface{}]bool)
	for _, val := range arr1 {
		if resultMap[val] {
			continue
		}
		resultMap[val] = true
	}

	log.Printf("obj: %+v\n", resultMap)

	for _, val2 := range arr2 {
		if resultMap[val2] {
			return true
		}
	}

	return false
}
