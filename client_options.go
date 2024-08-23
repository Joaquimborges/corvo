package corvo

type requestOptions func(*clientOptions)

type EndpointURL string

const (
	//	/token/v1/autentica/cartaopostagem
	GenerateAccessTokenURL       EndpointURL = "generate_game_token"
	CheckDeliveryDueDateURL      EndpointURL = "check_delivery_due_date"
	CheckDeliveryProductPriceURL EndpointURL = "check_delivery_product_price"
)

type clientOptions struct {
	body    any
	decode  any
	headers map[string]string
}

func withBody(body any) requestOptions {
	return func(co *clientOptions) {
		co.body = body
	}
}

func withHeaders(headers map[string]string) requestOptions {
	return func(co *clientOptions) {
		co.headers = headers
	}
}

func withDecodeValue(decode any) requestOptions {
	return func(co *clientOptions) {
		co.decode = decode
	}
}
