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
	FloatPrice float64 `json:"preco_float,omitempty"`
	StrPrice   string  `json:"preco_str,omitempty"`
}

func newDeliveryPrice(floatAmount float64, strAmount string) *DeliveryPrice {
	return &DeliveryPrice{
		FloatPrice: floatAmount,
		StrPrice:   strAmount,
	}
}
