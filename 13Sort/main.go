package main

import (
	"fmt"
	"mysort/sortfuncs"
)

func main() {
	arrForInsertionSort := []int{5, 2, 9, 1, 5, 6}
	fmt.Printf("Массив для сортировки простыми вставками: %v\n", arrForInsertionSort)
	sortedArrInsertionSort := sortfuncs.InsertionSort(arrForInsertionSort)
	fmt.Printf("Отсортированный простыми вставками массив: %v\n", sortedArrInsertionSort)

	arrForBubbleSort := []int{4, 5, 3, 1, 2, 6, 8, 10}
	fmt.Printf("Массив для сортировки пузырьками: %v\n", arrForBubbleSort)
	sortedArrForBubbleSort := sortfuncs.BubbleSort(arrForBubbleSort)
	fmt.Printf("Отсортированный пузырьками массив: %v\n", sortedArrForBubbleSort)

	arrForQuickSort := []int{7, 3, 1, 34, 12, 23, 4, 6}
	fmt.Printf("Массив для быстрой сортировки: %v\n", arrForQuickSort)
	sortedArrForQuickSort := sortfuncs.MyQuickSort(arrForQuickSort)
	fmt.Printf("Отсортированный быстрой сортировкой массив: %v\n", sortedArrForQuickSort)
}
