package currency

const (
	RUB = "RUB"
	USD = "USD"
	EUR = "EUR"
)

type Currency struct {
	Curr []string
}

func New() Currency {
	return Currency{
		Curr: []string{RUB, USD, EUR},
	}
}
