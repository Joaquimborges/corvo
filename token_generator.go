package corvo

import (
	"fmt"
	"net/http"
)

type tokenData struct {
	Environment string   `json:"ambiente"`
	Cnpj        string   `json:"cnpj"`
	PostCard    postCard `json:"cartaoPostagem"`
	Token       string   `json:"token"`
}

type postCard struct {
	Number   string `json:"numero"`
	Contract string `json:"contrato"`
}

type requestBody struct {
	PostCardNumber string `json:"numero"`
}

func generateAccessToken(config *Config, httpClient *restClient) (*tokenData, error) {
	headers := map[string]string{
		"authorization": fmt.Sprintf("Basic %s", config.AuthorizationCode),
	}

	body := requestBody{PostCardNumber: config.PostCard}
	var responseData tokenData

	err := httpClient.BuildRequest(
		config.UrlMapper[GenerateAccessTokenURL],
		http.MethodPost,
		WithBody(body),
		WithHeaders(headers),
		WithDecodeValue(&responseData),
	).Execute()

	if err != nil {
		return nil, err
	}
	return &responseData, nil
}
