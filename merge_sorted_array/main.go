package main

import (
	"log"
)

func main() {
	arr1 := []int{1, 3, 5}
	arr2 := []int{0, 2, 4, 6, 7, 8}

	log.Println(mergeSortedArray(arr1, arr2))
}

func mergeSortedArray(arr1, arr2 []int) []int {
	var mergeArr []int
	arr1Idx, arr2Idx := 0, 0

	for len(arr1) > arr1Idx && len(arr2) > arr2Idx {
		if arr1[arr1Idx] < arr2[arr2Idx] {
			mergeArr = append(mergeArr, arr1[arr1Idx])
			arr1Idx++
		} else {
			mergeArr = append(mergeArr, arr2[arr2Idx])
			arr2Idx++
		}
	}

	// if there is any left items in arr1, push them to mergeArr
	for len(arr1) > arr1Idx {
		mergeArr = append(mergeArr, arr1[arr1Idx])
		arr1Idx++
	}

	// if there is any left items in arr2, push them to mergeArr
	for len(arr2) > arr2Idx {
		mergeArr = append(mergeArr, arr2[arr2Idx])
		arr2Idx++
	}

	return mergeArr
}
