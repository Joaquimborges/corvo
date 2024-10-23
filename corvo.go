package corvo

import (
	"errors"
	"fmt"
	"net/http"
)

type WebServices interface {
	CheckDeliveryDueDate(productCode string, destineZipCode string) (*DeliveryTimeResponse, error)
	CheckDeliveryProductPrice(productCode string, destineZipCode string, dimensions *ProductDimensions) (*DeliveryPrice, error)
}

type webServices struct {
	client *restClient
	config *Config
}

func NewCorreiosWebServices(config *Config) WebServices {
	return &webServices{
		client: newHttpClient(),
		config: config,
	}
}

func (ws *webServices) CheckDeliveryDueDate(productCode string, destineZipCode string) (*DeliveryTimeResponse, error) {
	headers, err := ws.buildRequestHeaders()
	if err != nil {
		return nil, fmt.Errorf("[CheckDeliveryDueDate] erro ao gerar o token de acesso: %v", err)
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
		withHeaders(headers),
		withDecodeValue(&response),
	).Execute()

	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (ws *webServices) CheckDeliveryProductPrice(productCode string, destineZipCode string, dimensions *ProductDimensions) (*DeliveryPrice, error) {
	headers, err := ws.buildRequestHeaders()
	if err != nil {
		return nil, fmt.Errorf("[CheckDeliveryProductPrice] erro ao gerar o token de acesso %v", err)
	}

	var requestURL string
	if url, ok := ws.config.UrlMapper[CheckDeliveryProductPriceURL]; !ok {
		return nil, errors.New("CheckDeliveryProductPriceURL não foi encontrada")
	} else {
		requestURL = url
	}

	if ws.config.useConfigDimensions {
		dimensions = ws.config.Dimensions
	}

	requestURL += fmt.Sprintf(
		"/%s?cepDestino=%s&cepOrigem=%s&psObjeto=%d&tpObjeto=%d&comprimento=%d&largura=%d&altura=%d&vlDeclarado=%d",
		productCode,
		destineZipCode,
		ws.config.OriginZipCode,
		dimensions.Weight,
		ws.config.DeliveryType,
		dimensions.Fulfillment,
		dimensions.Height,
		dimensions.Width,
		ws.config.DefaultDeclaredValue,
	)

	for _, additionalService := range ws.config.AdditionalServices {
		requestURL += fmt.Sprintf("&servicosAdicionais=%s", additionalService)
	}

	var response DeliveryPriceResponse
	err = ws.client.BuildRequest(
		requestURL,
		http.MethodGet,
		withHeaders(headers),
		withDecodeValue(&response),
	).Execute()

	if err != nil {
		return nil, err
	}

	if response.FinalPrice != "" {
		if ws.config.shouldGenerateFloatPrice {
			price, err := parseBrazilianStrAmountToFloat(response.FinalPrice)
			if err != nil {
				return nil, fmt.Errorf("erro ao converter o preço de string para float: %v", err)
			}
			return newDeliveryPrice(price, ""), nil
		}
		return newDeliveryPrice(0, response.FinalPrice), nil
	}
	return nil, errors.New("a requisição provavelmente retornou erro")
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
