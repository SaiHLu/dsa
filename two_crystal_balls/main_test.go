package main

import (
	"log"
	"testing"
)

func TestBreakingPoints(t *testing.T) {
	breaks := []bool{
		false, false, false,
		false, false, false,
		false, false, true,
		true, true, true,
		true, true, true,
		true,
	}

	result := breakingPoints(breaks)

	if result == -1 {
		t.Errorf("Failed: return %v", result)
	}

	log.Println("result: ", result)
}
