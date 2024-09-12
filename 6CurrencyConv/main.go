package main

import (
	"fmt"
)

const (
	RUB = "RUB"
	USD = "USD"
	EUR = "EUR"
)

var CURRENCY = []string{RUB, USD, EUR}

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
	{USD, EUR, 0.85},
	{EUR, USD, 1.18},
	{USD, RUB, 74.0},
	{RUB, USD, 0.0135},
	{EUR, RUB, 87.0},
	{RUB, EUR, 0.0115},
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

func readEx() ExchangeRate {
	var from string                        // переменная для хранения исходной валюты
	var to string                          // переменная для хранения целевой валюты
	var amount float64                     // переменная для хранения суммы конвертации
	fmt.Print("Введите исходную валюту: ") // in CURRENCY
	fmt.Scan(&from)
	fmt.Print("Введите целевую валюту: ") // in CURRENCY
	fmt.Scan(&to)
	fmt.Print("Введите сумму для конвертации: ") // >0 ant digit
	fmt.Scan(&amount)
	ex := ExchangeRate{
		From:   from,
		To:     to,
		Amount: amount,
	}
	return ex
}

func main() {
	for {
		ex := readEx()
		converted, err := exchange(&ex)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(converted)
		}
	}

}
