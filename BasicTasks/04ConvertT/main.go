package main

import (
	"fmt"
	"strconv"
)

func celsToFahr(tempString string) float64 {
	var tempCelsInt int // переменная в которую мы сохраним введенную строку в интах
	tempCelsInt, err := strconv.Atoi(tempString)
	if err != nil {
		panic(err)
	}
	var tempFahr float64 // переменная в которую сохраним температуру в фаренгейтах
	tempFahr = 1.8*float64(tempCelsInt) + 32
	return tempFahr
}
func main() {
	var tempC string //переменнаая в которую сохраняем результат ввода пользователя
	for {
		fmt.Print("Введите температуру в градусах цельсия (end - завершение програмы): ")
		_, err := fmt.Scanln(&tempC)
		if err != nil {
			fmt.Printf("Неподходящее значение: %s\n", err.Error())
		} else if tempC == "end" {
			break
		} else {
			fmt.Println(celsToFahr(tempC))
		}
	}
}
