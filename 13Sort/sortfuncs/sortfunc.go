package sortfuncs

// 1. Сортировка пузырьком

// 2. Сортировка простыми вставками
func InsertionSort(arr []int) []int { // O(N^2)
	for i := 1; i < len(arr); i++ { // O(n)
		key := arr[i]
		j := i - 1

		// Перемещаем элементы arr[0..i-1], которые больше key,
		// на одну позицию вперед от их текущей позиции
		for j >= 0 && arr[j] > key { // O(n)
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

// 3. qsort (быстрая сорировка)
func MyQuickSort(arr []int) []int {
	return arr
}
