package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s)
	stringLength := len(runes) - 1
	var reverseStr []rune

	for i := stringLength; i >= 0; i-- {
		reverseStr = append(reverseStr, runes[i])
	}

	return string(reverseStr)
}

func main() {
	str := "Hello world"
	reversedStr := reverseString(str)

	fmt.Println(reversedStr)
}
