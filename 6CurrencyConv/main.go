package main

import (
	"currencyconv/convertor"
	"fmt"
	"strings"
)

func readEx() (string, string, float64, error) {
	var from string    // переменная для хранения исходной валюты
	var to string      // переменная для хранения целевой валюты
	var amount float64 // переменная для хранения суммы конвертации
	curr := convertor.New()
	for {
		fmt.Print("Введите исходную валюту: ")
		_, err := fmt.Scan(&from)
		if err != nil {
			return from, to, amount, err
		}
		from = strings.ToUpper(from)

		if curr.IsCorrectCurr(from) {
			fmt.Println("Используйте другую валюту! (rub, usd, eur)")
		} else {
			break
		}
	}

	for {
		fmt.Print("Введите требуемую валюту: ")
		_, err := fmt.Scan(&to)
		if err != nil {
			return from, to, amount, err
		}
		to = strings.ToUpper(to)

		if curr.IsCorrectCurr(to) {
			fmt.Println("Используйте другую валюту! (rub, usd, eur)")
		} else {
			break
		}
	}
	for {
		fmt.Print("Введите значение для перевода: ")
		if _, err := fmt.Scan(&amount); err != nil {
			return from, to, amount, err
		}

		if amount < 0 {
			fmt.Println("Введите число больше нуля")
		} else {
			break
		}
	}

	return from, to, amount, nil // Чтобы функция main завершила выполнение было errors.New("Моя ошибка!")
}

func main() {
	conv := convertor.New() // пакет convertor, функция new
	for {

		from, to, amount, err := readEx()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		converted, err := conv.Convert(from, to, amount)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(converted)
		}
	}

}
