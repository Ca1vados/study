package entity

type ApiBinanceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type Responce struct {
	FromTo string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
}
