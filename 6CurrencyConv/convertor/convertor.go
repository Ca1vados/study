package convertor

import (
	"errors"
	"fmt"
)

const (
	RUB = "RUB"
	USD = "USD"
	EUR = "EUR"
)

type Convertor struct {
	Curr  []string
	Ratio map[string]float64
}

func New() Convertor {
	ratio := map[string]float64{
		USD + "-" + EUR: 0.85,
		EUR + "-" + USD: 1.18,
		USD + "-" + RUB: 74.0,
		RUB + "-" + USD: 0.0135,
		EUR + "-" + RUB: 87.0,
		RUB + "-" + EUR: 0.0115,
	}

	return Convertor{
		Curr:  []string{RUB, USD, EUR},
		Ratio: ratio,
	}
}

func (c *Convertor) IsCorrectCurr(curr string) bool {
	for _, v := range c.Curr {
		if v == curr {
			return true
		}
	}

	return false
}

func (c *Convertor) Convert(from string, to string, amount float64) (float64, error) {
	key := from + "-" + to

	if !c.IsCorrectCurr(from) {
		return 0, errors.New("Value 'from' is incorrect ...")
	}

	if !c.IsCorrectCurr(to) {
		return 0, errors.New("Value 'to' is incorrect ...")
	}

	k, ok := c.Ratio[key] // k -значение по ключу, ok - будет true, если есть такой ключ
	if !ok {
		return 0, errors.New(fmt.Sprintf("Key %s is incorrect ...", key))
	}

	res := amount * k

	return res, nil
}
