package corvo

const (
	PackServiceCode = "03310"
)

type DeliveryTimeResponse struct {
	DeliveryTime uint64 `json:"prazoEntrega"`
	MaxDueDate   string `json:"dataMaxima"`
}
