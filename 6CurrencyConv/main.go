package main

import (
	"fmt"
)

type ExchangeRate struct {
	From string
	To   string
	Rate float64
}

var rates = []ExchangeRate{
	{"USD", "EUR", 0.85},
	{"EUR", "USD", 1.18},
	{"USD", "RUB", 74.0},
	{"RUB", "USD", 0.0135},
	{"EUR", "RUB", 87.0},
	{"RUB", "EUR", 0.0115},
}

func exchange(f, t string, am float64) {
	for _, i := range rates {
		if rates.From == f || rates.To == t {
			return am * rates.Rate
		}
		fmt.Println(c)
	}
}

func main() {

	var from string    // переменная для хранения исходной валюты
	var to string      // переменная для хранения целевой валюты
	var amount float64 // переменная для хранения суммы конвертации

	fmt.Print("Введите исходную валюту: ")
	fmt.Scan(&from)
	fmt.Print("Введите целевую валюту: ")
	fmt.Scan(&to)
	fmt.Print("Введите сумму для конвертации: ")
	fmt.Scan(&amount)
	exchange(from, to, amount)
}
