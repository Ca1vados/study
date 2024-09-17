package main

import (
	"currencyconv/currency"
	"errors"
	"fmt"
)

type Convertor struct {
	From string
	To   string
	Rate float64
}

type ExchangeRate struct {
	From   string
	To     string
	Amount float64
}

var rates = []Convertor{
	{currency.USD, currency.EUR, 0.85},
	{currency.EUR, currency.USD, 1.18},
	{currency.USD, currency.RUB, 74.0},
	{currency.RUB, currency.USD, 0.0135},
	{currency.EUR, currency.RUB, 87.0},
	{currency.RUB, currency.EUR, 0.0115},
}

func exchange(er *ExchangeRate) (float64, error) {
	var res float64 = 0.0
	convertOk := false
	for _, con := range rates {
		if con.From == er.From && con.To == er.To {
			res = er.Amount * con.Rate
			convertOk = true
		}
	}
	if convertOk {
		return res, nil
	} else {
		return res, fmt.Errorf("Bad convert :(")
	}

}

func readEx() (ExchangeRate, error) {
	var from string // переменная для хранения исходной валюты
	//var to string      // переменная для хранения целевой валюты
	//var amount float64 // переменная для хранения суммы конвертации
	curr := currency.New()
	var ex ExchangeRate
	for {
		fmt.Print("Введите исходную валюту: ")
		fmt.Scan(&from)
		f := false
		for _, v := range curr.Curr {
			if v == from {
				f = true
				break
			}
		}

		if !f {
			fmt.Println("Используйте другую валюту!")
		} else {
			break
		}
	}

	return ex, errors.New("Моя ошибка!") // Чтобы функция main завершила выполнение
}

func main() {
	for {
		ex, err := readEx()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		converted, err := exchange(&ex)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(converted)
		}
	}

}
