package main

import (
	"fmt"
	"mysort/sortfuncs"
)

func main() {
	arr := []int{5, 2, 9, 1, 5, 6}
	sortedArr := sortfuncs.InsertionSort(arr)
	fmt.Println(sortedArr)
}
