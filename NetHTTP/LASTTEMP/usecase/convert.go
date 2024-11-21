package usecase

import (
	"fmt"
	"strconv"
	"temp/entity"
)

func getCoefficient(arr []entity.ApiBinanceResponse, symbol string) (float64, error) {
	var coefString string
	var f = false
	for _, i := range arr {
		if i.Symbol == symbol {
			coefString = i.Price
			f = true
			break
		}
	}
	if f == false {
		return 1, fmt.Errorf("Нет варианта перевода: %v", symbol)
	}
	coefFloat64, err := strconv.ParseFloat(coefString, 64)
	if err != nil {
		return 1, fmt.Errorf("Проблемы с переводом Price в float64: %v", err)
	}
	return coefFloat64, nil
}

func (u *UseCase) Convert(req entity.ConvertRequest) (entity.ConvertResponse, error) {
	res := entity.ConvertResponse{
		Amount: req.Amount,
		From:   req.From,
		To:     req.To,
	}
	cryptoRates, err := u.GetCryptoRatesBinance()
	if err != nil {
		return res, fmt.Errorf("нет данных в сети: %v", err)
	}

	if req.From == "BTC" {
		if req.To == "USDT" {
			symbol := req.From + req.To
			k, _ := getCoefficient(cryptoRates, symbol)
			res.Result = res.Amount / k
		} else if req.To == "ETH" {
			symbol1 := req.From + "USDT"
			k1, _ := getCoefficient(cryptoRates, symbol1)
			symbol2 := req.To + "USDT"
			k2, _ := getCoefficient(cryptoRates, symbol2)
			res.Result = res.Amount / k1 * k2
		} else if req.To == "LTC" {
			symbol1 := req.From + "USDT"
			k1, _ := getCoefficient(cryptoRates, symbol1)
			symbol2 := req.To + "USDT"
			k2, _ := getCoefficient(cryptoRates, symbol2)
			res.Result = res.Amount / k1 * k2
		} else {
			return res, fmt.Errorf("неизвестная валюта %v", req.To)
		}
	} else if req.From == "ETH" {
		if req.To == "USDT" {
			symbol := req.From + req.To
			k, _ := getCoefficient(cryptoRates, symbol)
			res.Result = res.Amount / k
		} else if req.To == "BTC" {
			symbol1 := req.From + "USDT"
			k1, _ := getCoefficient(cryptoRates, symbol1)
			symbol2 := req.To + "USDT"
			k2, _ := getCoefficient(cryptoRates, symbol2)
			res.Result = res.Amount / k1 * k2
		} else if req.To == "LTC" {
			symbol1 := req.From + "USDT"
			k1, _ := getCoefficient(cryptoRates, symbol1)
			symbol2 := req.To + "USDT"
			k2, _ := getCoefficient(cryptoRates, symbol2)
			res.Result = res.Amount / k1 * k2
		} else {
			return res, fmt.Errorf("неизвестная валюта %v", req.To)
		}
	} else if req.From == "LTC" {
		if req.To == "USDT" {
			symbol := req.From + req.To
			k, _ := getCoefficient(cryptoRates, symbol)
			res.Result = res.Amount / k
		} else if req.To == "ETH" {
			symbol1 := req.From + "USDT"
			k1, _ := getCoefficient(cryptoRates, symbol1)
			symbol2 := req.To + "USDT"
			k2, _ := getCoefficient(cryptoRates, symbol2)
			res.Result = res.Amount / k1 * k2
		} else if req.To == "BTC" {
			symbol1 := req.From + "USDT"
			k1, _ := getCoefficient(cryptoRates, symbol1)
			symbol2 := req.To + "USDT"
			k2, _ := getCoefficient(cryptoRates, symbol2)
			res.Result = res.Amount / k1 * k2
		} else {
			return res, fmt.Errorf("неизвестная валюта %v", req.To)
		}
	} else if req.From == "USDT" {
		if req.To == "LTC" {
			symbol := req.To + req.From
			k, _ := getCoefficient(cryptoRates, symbol)
			res.Result = res.Amount * k
		} else if req.To == "ETH" {
			symbol := req.To + req.From
			k, _ := getCoefficient(cryptoRates, symbol)
			res.Result = res.Amount * k
		} else if req.To == "BTC" {
			symbol := req.To + req.From
			k, _ := getCoefficient(cryptoRates, symbol)
			res.Result = res.Amount * k
		} else {
			return res, fmt.Errorf("неизвестная валюта %v", req.To)
		}
	} else {
		return res, fmt.Errorf("неизвестная валюта %v", req.From)
	}
	return res, nil
}
