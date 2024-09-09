package main

import (
	"fmt"
	"math"
)

// функция, определяющая, является ли число простым
func isItPrimeNumber(n int) bool {
	if n == 0 || n == 1 {
		return false
	} else {
		for i := 2; i <= n/2; i++ {
			if n%i == 0 {
				return false
			}
		}
	}
	return true
}

// функция, округляющая введенное число вверх и выводящая слайс простых чисел до введенного числа
func outputSlicePrimeNumbers(num float64) {
	roundedNumber := int(math.Ceil(num)) // переменная в которую помещаем округленное вверх число
	numbers := make([]int, 0)            // создаем слайс для простых чисел
	for i := 2; i < roundedNumber; i++ {
		if isItPrimeNumber(i) == true {
			numbers = append(numbers, i)
		}
	}
	fmt.Println(numbers)
}

func main() {
	var number float64 //переменнаая в которую сохраняем результат ввода пользователя. На всякий пожарный float64, если пользователь захочет ввести число с запятой
	fmt.Print("Введите число: ")
	fmt.Scanln(&number)
	outputSlicePrimeNumbers(number)
}
