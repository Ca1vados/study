package main

import (
	"testing"
)

func TestSwap(t *testing.T) {
	resA := 2
	resB := 3
	Swap(&resA, &resB)
	expectedA := 3
	expectedB := 2
	if resA != expectedA || resB != expectedB {
		t.Errorf("Swap(2, 3) = %d, %d; want %d, %d", resA, resB, expectedA, expectedB)
	}
}
