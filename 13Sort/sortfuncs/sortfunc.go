package sortfuncs

import "math/rand"

// 1. Сортировка пузырьком
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ { // загоняем все элементы слайса в цикл
		for j := 0; j < len(arr)-1-i; j++ { // загоняем во второй цикл все елементы слайса минус по сути количество предыдущих циклов
			if arr[j] > arr[j+1] { // сравниваем  попарно все елементы
				arr[j], arr[j+1] = arr[j+1], arr[j] // если первый элемент больше второго, меняем местами
			}
		}
	}
	return arr
}

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
	if len(arr) <= 1 { // чтобы избежать бесконечной рекурсии
		return arr
	}

	pivot := arr[rand.Intn(len(arr))] // найдем опорное число(a-ka median, pivot) случайным выбором, можно применять разные алгоритмы выбора опорного числа

	lowPart := make([]int, 0, len(arr))    // в этот слайс мы запишем числа меньше опорного числа
	highPart := make([]int, 0, len(arr))   // в этот слайс мы запишем числа больше опорного числа
	middlePart := make([]int, 0, len(arr)) // в этот слайс мы запишем числа равные опорному числу

	for _, i := range arr { // перебираем массив arr
		switch {
		case i < pivot:
			lowPart = append(lowPart, i) // если число меньше опорного, добавляем его в слайс lowPart
		case i == pivot:
			middlePart = append(middlePart, i) // если число равно опорному, добавляем его в слайс middlePart
		case i > pivot:
			highPart = append(highPart, i) // если число больше опорного, добавляем его в слайс highPart
		}
	}

	lowPart = MyQuickSort(lowPart)   // рекурсивно вызываем функцию QuickSort для слайса меньших чисел
	highPart = MyQuickSort(highPart) // рекурсивно вызываем функцию QuickSort для слайса больших чисел

	// собираем итоговый слайс

	arr = append(lowPart, middlePart...) // добавляем в наш слайс lowPart и middlePart
	arr = append(arr, highPart...)       // дабавляем в наш слайс highPart
	return arr
}
