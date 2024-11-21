package entity

type ApiBinanceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type ConvertRequest struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

type ConvertResponse struct {
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	Result float64 `json:"result"`
	To     string  `json:"to"`
}
