package corvo

import "time"

type RequestOptions func(*clientOptions)

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
	timeout time.Duration
}

func WithBody(body any) RequestOptions {
	return func(co *clientOptions) {
		co.body = body
	}
}

func WithHeaders(headers map[string]string) RequestOptions {
	return func(co *clientOptions) {
		co.headers = headers
	}
}

func WithDecodeValue(decode any) RequestOptions {
	return func(co *clientOptions) {
		co.decode = decode
	}
}

func WithTimeout(timeout time.Duration) RequestOptions {
	return func(co *clientOptions) {
		co.timeout = timeout
	}
}
