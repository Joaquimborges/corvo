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

	returnErrorIf := func(req *http.Request) bool {
		destineZipCode := req.URL.Query().Get("cepDestino")
		return len(destineZipCode) < 8
	}
	errorResponseBody := []byte(`{"message":"invalid destine zip code"}`)

	server := httptest.NewServer(mocks.BuildTestHandleFunc(
		returnErrorIf,
		errorResponseBody,
		buildDeliveryDueDateBytes(),
		http.StatusOK,
		http.StatusBadRequest,
		t,
	))

	defer tokenServer.Close()
	defer server.Close()

	t.Run("should return error on generate access token", func(t *testing.T) {
		config := &Config{
			PostCard:          "00112233",
			AuthorizationCode: "foo bar",
			UrlMapper: map[urlKey]string{
				GenerateAccessTokenUrlKey:  tokenServer.URL,
				CheckDeliveryDueDateUrlKey: server.URL,
			},
		}

		wServices := NewCorreiosWebServices(config)
		response, err := wServices.CheckDeliveryDueDate("03310", "05746000", "44320000")

		require.Error(t, err)
		require.Nil(t, response)
		require.Contains(t, err.Error(), "bad credential")
	})

	t.Run("should return error, invalid destine zip code", func(t *testing.T) {
		config := &Config{
			PostCard:          "00112233",
			AuthorizationCode: "foo",
			UrlMapper: map[urlKey]string{
				GenerateAccessTokenUrlKey:  tokenServer.URL,
				CheckDeliveryDueDateUrlKey: server.URL,
			},
		}

		wServices := NewCorreiosWebServices(config)
		response, err := wServices.CheckDeliveryDueDate("03310", "05746000", "00")

		require.Error(t, err)
		require.Nil(t, response)
		require.Contains(t, err.Error(), "invalid destine zip code")
	})

	t.Run("should return ok", func(t *testing.T) {
		config := &Config{
			PostCard:          "00112233",
			AuthorizationCode: "foo",
			UrlMapper: map[urlKey]string{
				GenerateAccessTokenUrlKey:  tokenServer.URL,
				CheckDeliveryDueDateUrlKey: server.URL,
			},
		}

		wServices := NewCorreiosWebServices(config)
		response, err := wServices.CheckDeliveryDueDate("03310", "05746000", "44320000")

		require.NoError(t, err)
		require.NotNil(t, response)
	})
}

func buildDeliveryDueDateBytes() []byte {
	return []byte(`{"coProduto": "03310","prazoEntrega": 10,"dataMaxima": "2024-09-02T23:59:59"}`)
}
