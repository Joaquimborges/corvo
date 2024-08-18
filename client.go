package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const BaseURL = "https://api.correios.com.br"

type restClient struct {
	httpClient     *http.Client
	requestOptions *clientOptions
	url            string
	method         string
}

func NewHttpClient() *restClient {
	return &restClient{}
}

func (client *restClient) BuildRequest(url, method string, options ...RequestOptions) *restClient {
	var requestOptions clientOptions
	for _, option := range options {
		option(&requestOptions)
	}

	httpClient := &http.Client{Timeout: 5 * time.Second}
	if requestOptions.timeout.String() != "0s" {
		httpClient.Timeout = requestOptions.timeout
	}

	client.httpClient = httpClient
	client.requestOptions = &requestOptions
	client.method = method
	client.url = url

	return client
}

func (client *restClient) Execute() error {
	var buf bytes.Buffer
	if client.requestOptions.body != nil {
		if err := json.NewEncoder(&buf).Encode(client.requestOptions.body); err != nil {
			return fmt.Errorf("error enconde body, message: %s", err.Error())
		}
	}

	request, err := http.NewRequest(client.method, client.url, &buf)
	if err != nil {
		return fmt.Errorf(
			"error build %s rest request, message: %v",
			client.method,
			err,
		)
	}

	for k, v := range client.requestOptions.headers {
		request.Header.Set(k, v)
	}

	data, er := client.doRequest(request)
	if er != nil {
		return er
	}

	if client.requestOptions.decode != nil {
		if err = json.Unmarshal(data, client.requestOptions.decode); err != nil {
			return fmt.Errorf("error Unmarshal response: %v", err)
		}
	}
	return nil
}

func (client *restClient) doRequest(req *http.Request) ([]byte, error) {
	resp, er := client.httpClient.Do(req)
	if er != nil {
		return nil, fmt.Errorf(
			"error doing client request, message: %s, "+
				"url: %s",
			er.Error(), req.URL.Path,
		)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		bytes, err := client.closeBodyAndSendResponse(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading body response: %v", err)
		}
		return nil, errors.New(string(bytes))
	}
	return client.closeBodyAndSendResponse(resp.Body)
}

func (client *restClient) closeBodyAndSendResponse(body io.ReadCloser) ([]byte, error) {
	bts, ioErr := io.ReadAll(body)
	if ioErr != nil {
		return nil, fmt.Errorf("error reading body response")
	}
	return bts, nil
}
