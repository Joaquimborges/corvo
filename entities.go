package corvo

const (
	PackServiceCode = "03310"
)

type DeliveryTimeResponse struct {
	DeliveryTime uint64 `json:"prazoEntrega"`
	MaxDueDate   string `json:"dataMaxima"`
}

type DeliveryPriceResponse struct {
	FinalPrice string `json:"pcFinal"`
}

type DeliveryPrice struct {
	Price float64 `json:"preco"`
}

func newDeliveryPrice(amount float64) *DeliveryPrice {
	return &DeliveryPrice{Price: amount}
}
