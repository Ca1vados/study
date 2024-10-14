package main

import (
	"fmt"
)

func DoubleArray(arr *[5]int) *[5]int {
	for a, _ := range arr {
		arr[a] *= 2
	}
	return arr
}
func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	DoubleArray(&numbers)
	fmt.Println(numbers)
}
