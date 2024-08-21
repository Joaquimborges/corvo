package corvo

import (
	"errors"
	"fmt"
	"net/http"
)

type WebServices interface {
	CheckDeliveryDueDate(productCode string, destineZipCode string) (*DeliveryTimeResponse, error)
	CheckDeliveryProductPrice(productCode string, destineZipCode string) (*DeliveryPrice, error)
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

func (ws *webServices) CheckDeliveryDueDate(productCode string, destineZipCode string) (*DeliveryTimeResponse, error) {
	headers, err := ws.buildRequestHeaders()
	if err != nil {
		return nil, fmt.Errorf("[CheckDeliveryDueDate] error on generateAccessToken: %v", err)
	}

	url := ws.config.UrlMapper[CheckDeliveryDueDateURL]
	url += fmt.Sprintf(
		"/%s?cepOrigem=%s&cepDestino=%s",
		productCode,
		ws.config.OriginZipCode,
		destineZipCode,
	)

	var response DeliveryTimeResponse
	err = ws.client.BuildRequest(
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

func (ws *webServices) CheckDeliveryProductPrice(productCode string, destineZipCode string) (*DeliveryPrice, error) {
	headers, err := ws.buildRequestHeaders()
	if err != nil {
		return nil, fmt.Errorf("[CheckDeliveryProductPrice] error on generateAccessToken: %v", err)
	}

	var requestURL string
	if url, ok := ws.config.UrlMapper[CheckDeliveryProductPriceURL]; !ok {
		return nil, errors.New("CheckDeliveryProductPriceURL was not founded")
	} else {
		requestURL = url
	}

	requestURL += fmt.Sprintf(
		"/%s?cepDestino=%s&cepOrigem=%s&psObjeto=%d&tpObjeto=%d&comprimento=%d&largura=%d&altura=%dvlDeclarado=%d",
		productCode,
		destineZipCode,
		ws.config.OriginZipCode,
		ws.config.ObjectBaseWeight,
		ws.config.DeliveryType,
		ws.config.BaseFulfillment,
		ws.config.BaseHeight,
		ws.config.BaseWidth,
		ws.config.DefaultDeclaredValue,
	)

	for _, additionalService := range ws.config.AdditionalServices {
		requestURL += fmt.Sprintf("&servicosAdicionais=%s", additionalService)
	}

	var response DeliveryPriceResponse
	err = ws.client.BuildRequest(
		requestURL,
		http.MethodGet,
		WithHeaders(headers),
		WithDecodeValue(&response),
	).Execute()

	if err != nil {
		return nil, err
	}

	if response.FinalPrice != "" {
		price, err := parseBrazilianStrAmountToFloat(response.FinalPrice)
		if err != nil {
			return nil, fmt.Errorf("error parsing brazilian string amount to float: %v", err)
		}
		return newDeliveryPrice(price), nil
	}
	return nil, errors.New("the external call may have returned an error")
}

func (ws *webServices) buildRequestHeaders() (map[string]string, error) {
	tokenData, err := generateAccessToken(ws.config, ws.client)
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("Bearer %s", tokenData.Token)
	return headers, nil
}
