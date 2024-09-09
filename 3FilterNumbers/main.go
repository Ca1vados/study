package main

import (
	"fmt"
	"strconv"
	"strings"
)

// функция, собирающая введенные через запятую значения в слайс строк
func makeSliceNumbers(str string) []string {
	slice := strings.Split(str, ",")
	return slice
}

// функция выводящая на печать слайс четных значений
func makeListEven(str string) []int {
	sliceInt := make([]int, 0) // создаем слайс, куда поместим конвертированные значения из нашего слайса строк
	for _, i := range makeSliceNumbers(str) {
		i, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		sliceInt = append(sliceInt, i)
	}
	sliceIntEven := make([]int, 0) // создаем слайс, куда поместим четные значения из нашего слайса интов
	for _, i := range sliceInt {
		if i%2 == 0 {
			sliceIntEven = append(sliceIntEven, i)
		}
	}
	return sliceIntEven
}

func main() {
	var line string
	fmt.Print("Напишите числа через запятую: ")
	fmt.Scan(&line)
	fmt.Println(makeListEven(line))
}
