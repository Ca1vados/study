package main

import (
	"mysort/sortfuncs"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 2, 9, 1, 5, 6}, []int{1, 2, 5, 5, 6, 9}},
		{[]int{3, 3, 3}, []int{3, 3, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{4, 2, 4, 54, 23, 12, 43, 54, 65, 76, 12, 433, 16, 11, 1, 87, 65, 43, 54, 32, 43, 12, 32, 54, 32, 12, 65, 32, 21, 54}, []int{1, 2, 4, 4, 11, 12, 12, 12, 12, 16, 21, 23, 32, 32, 32, 32, 43, 43, 43, 54, 54, 54, 54, 54, 65, 65, 65, 76, 87, 433}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := sortfuncs.InsertionSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 2, 9, 1, 5, 6}, []int{1, 2, 5, 5, 6, 9}},
		{[]int{3, 3, 3}, []int{3, 3, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{4, 2, 4, 54, 23, 12, 43, 54, 65, 76, 12, 433, 16, 11, 1, 87, 65, 43, 54, 32, 43, 12, 32, 54, 32, 12, 65, 32, 21, 54}, []int{1, 2, 4, 4, 11, 12, 12, 12, 12, 16, 21, 23, 32, 32, 32, 32, 43, 43, 43, 54, 54, 54, 54, 54, 65, 65, 65, 76, 87, 433}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := sortfuncs.BubbleSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 2, 9, 1, 5, 6}, []int{1, 2, 5, 5, 6, 9}},
		{[]int{3, 3, 3}, []int{3, 3, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{4, 2, 4, 54, 23, 12, 43, 54, 65, 76, 12, 433, 16, 11, 1, 87, 65, 43, 54, 32, 43, 12, 32, 54, 32, 12, 65, 32, 21, 54}, []int{1, 2, 4, 4, 11, 12, 12, 12, 12, 16, 21, 23, 32, 32, 32, 32, 43, 43, 43, 54, 54, 54, 54, 54, 65, 65, 65, 76, 87, 433}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := sortfuncs.MyQuickSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
