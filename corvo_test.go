package corvo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Joaquimborges/corvo/mocks"
	"github.com/stretchr/testify/require"
)

func TestCheckDeliveryDueDate(t *testing.T) {
	tokenServer := mocks.BuildGenerateAccessTokenTestSever(t)
	errorResponseBody := []byte(`{"message":"invalid destine zip code"}`)

	server := httptest.NewServer(mocks.BuildTestHandleFunc(
		mocks.DefaultDestineZipCodeRequestValidation,
		errorResponseBody,
		buildDeliveryDueDateBytes(),
		http.StatusOK,
		http.StatusBadRequest,
		t,
	))

	defer tokenServer.Close()
	defer server.Close()

	t.Run("should return error on generate access token", func(t *testing.T) {
		config := buildConfigs(tokenServer.URL, CheckDeliveryDueDateURL, server.URL)
		config.PostCard = "00112233"
		config.OriginZipCode = "44320000"
		config.AuthorizationCode = "foo bar"

		wServices := NewCorreiosWebServices(config)
		response, err := wServices.CheckDeliveryDueDate("03310", "05746000")

		require.Error(t, err)
		require.Nil(t, response)
		require.Contains(t, err.Error(), "bad credential")
	})

	t.Run("should return error, invalid destine zip code", func(t *testing.T) {
		config := buildConfigs(tokenServer.URL, CheckDeliveryDueDateURL, server.URL)
		config.PostCard = "00112233"
		config.OriginZipCode = "05746000"
		config.AuthorizationCode = "foo"

		wServices := NewCorreiosWebServices(config)
		response, err := wServices.CheckDeliveryDueDate("03310", "00")

		require.Error(t, err)
		require.Nil(t, response)
		require.Contains(t, err.Error(), "invalid destine zip code")
	})

	t.Run("should return ok", func(t *testing.T) {
		config := buildConfigs(tokenServer.URL, CheckDeliveryDueDateURL, server.URL)
		config.PostCard = "00112233"
		config.OriginZipCode = "44320000"
		config.AuthorizationCode = "foo"

		wServices := NewCorreiosWebServices(config)
		response, err := wServices.CheckDeliveryDueDate("03310", "05746000")

		require.NoError(t, err)
		require.NotNil(t, response)
	})
}

func TestCheckDeliveryProductPrice(t *testing.T) {
	tokenServer := mocks.BuildGenerateAccessTokenTestSever(t)
	errorResponseBody := []byte(`{"message":"invalid destine zip code"}`)

	server := httptest.NewServer(mocks.BuildTestHandleFunc(
		mocks.DefaultDestineZipCodeRequestValidation,
		errorResponseBody,
		buildDeliveryPriceBytes(),
		http.StatusOK,
		http.StatusBadRequest,
		t,
	))

	defer tokenServer.Close()
	defer server.Close()

	t.Run("should return error on generate access token", func(t *testing.T) {
		config := buildConfigs(tokenServer.URL, CheckDeliveryProductPriceURL, server.URL)
		config.PostCard = "00112233"
		config.OriginZipCode = "44320000"
		config.AuthorizationCode = "foo bar"

		wServices := NewCorreiosWebServices(config)
		response, err := wServices.CheckDeliveryProductPrice("03310", "05746000")

		require.Error(t, err)
		require.Nil(t, response)
		require.Contains(t, err.Error(), "bad credential")
	})

	t.Run("should return error, CheckDeliveryProductPriceURL not found", func(t *testing.T) {
		config := buildConfigs(tokenServer.URL, "", "")
		config.PostCard = "00112233"
		config.OriginZipCode = "44320000"
		config.AuthorizationCode = "foo"

		wServices := NewCorreiosWebServices(config)
		response, err := wServices.CheckDeliveryProductPrice("03310", "05746000")

		require.Error(t, err)
		require.Nil(t, response)
		require.Equal(t, "CheckDeliveryProductPriceURL was not founded", err.Error())
	})

	t.Run("should return error, invalid destine zip code", func(t *testing.T) {
		config := buildConfigs(tokenServer.URL, CheckDeliveryProductPriceURL, server.URL)
		config.PostCard = "00112233"
		config.OriginZipCode = "44320000"
		config.AuthorizationCode = "foo"
		config.DefaultDeclaredValue = 200
		config.ObjectBaseWeight = 400
		config.BaseFulfillment = 20
		config.BaseHeight = 20
		config.BaseWidth = 20
		config.DeliveryType = 2
		config.AdditionalServices = []string{"001", "019"}

		wServices := NewCorreiosWebServices(config)
		response, err := wServices.CheckDeliveryProductPrice("03310", "00")

		require.Error(t, err)
		require.Nil(t, response)
		require.Contains(t, err.Error(), "invalid destine zip code")
	})

	t.Run("should return error, invalid body response", func(t *testing.T) {
		server2 := httptest.NewServer(mocks.BuildTestHandleFunc(
			mocks.DefaultDestineZipCodeRequestValidation,
			errorResponseBody,
			[]byte(`{"coProduto": "03310","pcFinal": "33*40"}`),
			http.StatusOK,
			http.StatusBadRequest,
			t,
		))

		config := buildConfigs(tokenServer.URL, CheckDeliveryProductPriceURL, server2.URL)
		config.PostCard = "00112233"
		config.OriginZipCode = "44320000"
		config.AuthorizationCode = "foo"
		config.DefaultDeclaredValue = 200
		config.ObjectBaseWeight = 400
		config.BaseFulfillment = 20
		config.BaseHeight = 20
		config.BaseWidth = 20
		config.DeliveryType = 2
		config.AdditionalServices = []string{"001", "019"}

		wServices := NewCorreiosWebServices(config)
		response, err := wServices.CheckDeliveryProductPrice("03310", "05746000")

		require.Error(t, err)
		require.Nil(t, response)
		require.Contains(t, err.Error(), "error parsing brazilian string amount to float")
	})

	t.Run("should return 200-OK", func(t *testing.T) {
		config := buildConfigs(tokenServer.URL, CheckDeliveryProductPriceURL, server.URL)
		config.PostCard = "00112233"
		config.OriginZipCode = "44320000"
		config.AuthorizationCode = "foo"
		config.DefaultDeclaredValue = 200
		config.ObjectBaseWeight = 400
		config.BaseFulfillment = 20
		config.BaseHeight = 20
		config.BaseWidth = 20
		config.DeliveryType = 2
		config.AdditionalServices = []string{"001", "019"}

		wServices := NewCorreiosWebServices(config)
		response, err := wServices.CheckDeliveryProductPrice("03310", "05746000")

		require.NoError(t, err)
		require.NotNil(t, response)
		require.Equal(t, float64(35.40), response.Price)
	})
}

func buildDeliveryDueDateBytes() []byte {
	return []byte(`{"coProduto": "03310","prazoEntrega": 10,"dataMaxima": "2024-09-02T23:59:59"}`)
}

func buildDeliveryPriceBytes() []byte {
	return []byte(`{"coProduto": "03310","pcFinal": "35,40"}`)
}

func buildConfigs(tokenURL string, endpointUrl EndpointURL, serverUrl string) *Config {
	return &Config{
		UrlMapper: map[EndpointURL]string{
			GenerateAccessTokenURL: tokenURL,
			endpointUrl:            serverUrl,
		},
	}
}
