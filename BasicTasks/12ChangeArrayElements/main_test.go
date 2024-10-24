package main

import (
	"testing"
)

func TestArrayChange(t *testing.T) {
	numbers := [5]int{3, 4, 5, 6, 7}
	expectedNumbers := [5]int{6, 8, 10, 12, 14}
	DoubleArray(&numbers)
	for i, n := range numbers {
		if n != expectedNumbers[i] {
			t.Errorf("DoubleArray(&[5]int{3, 4, 5, 6, 7}) = %v; want %v", numbers, expectedNumbers)
		}
	}

}
