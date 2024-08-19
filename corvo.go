package main

import (
	"fmt"
	"net/http"
)

type WebServices interface {
	CheckDeliveryDueDate(product string, originZipCode string, destineZipCode string) (*DeliveryTimeResponse, error)
}

type webServices struct {
	client *restClient
	config *Config
}

func NewCorreiosWebServices(config *Config) WebServices {
	return &webServices{
		client: NewHttpClient(),
		config: config,
	}
}

func (service *webServices) CheckDeliveryDueDate(product string, originZipCode string, destineZipCode string) (*DeliveryTimeResponse, error) {
	headers, err := service.buildRequestHeaders()
	if err != nil {
		return nil, fmt.Errorf("[CheckDeliveryTime] error on generateAccessToken: %v", err)
	}

	url := service.config.UrlMapper[CheckDeliveryDueDateUrlKey]
	url += fmt.Sprintf("/%s?cepOrigem=%s&cepDestino=%s", product, originZipCode, destineZipCode)

	var response DeliveryTimeResponse

	err = service.client.BuildRequest(
		url,
		http.MethodGet,
		WithHeaders(headers),
		WithDecodeValue(&response),
	).Execute()

	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (service *webServices) buildRequestHeaders() (map[string]string, error) {
	tokenData, err := generateAccessToken(service.config, service.client)
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("Bearer %s", tokenData.Token)
	return headers, nil
}
