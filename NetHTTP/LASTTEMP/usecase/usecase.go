package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
	"temp/config"
	"temp/database"
	"temp/entity"
)

type UseCase struct {
	logLvl string
	db     *database.DataBase
}

func New(db *database.DataBase, cfg *config.Config) *UseCase {
	usecase := UseCase{
		logLvl: cfg.LogLvl,
		db:     db,
	}
	return &usecase
}

func (u *UseCase) GetGetCryptoRatesBinance() ([]entity.ApiBinanceResponse, error) {
	urls := []string{
		"https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT", //bitcoin(BTC) => ustd
		"https://api.binance.com/api/v3/ticker/price?symbol=ETHUSDT",
		"https://api.binance.com/api/v3/ticker/price?symbol=LTCUSDT",
	}
	result := make([]entity.ApiBinanceResponse, len(urls))

	for i, url := range urls {
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("1. %v\n", err.Error())
			return nil, err
		}
		defer response.Body.Close()

		// Проверка кода статуса ответа
		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error: received status code %d", response.StatusCode)
		}

		var conv_data entity.ApiBinanceResponse
		if err := json.NewDecoder(response.Body).Decode(&conv_data); err != nil {
			return nil, err
		}
		result[i] = conv_data
	}

	return result, nil
}
