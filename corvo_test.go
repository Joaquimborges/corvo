package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckDeliveryDueDate(t *testing.T) {
	tokenServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if authorization := r.Header.Get("authorization"); authorization != "Basic foo" {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte(`{"message":"bad credential"}`))
			if err != nil {
				t.Errorf("error write the response body: %v", err)
			}
		} else {
			w.WriteHeader(http.StatusCreated)
			responseBody := buildTokenResponseBytes(t, buildTokenData())
			_, err := w.Write(responseBody)
			if err != nil {
				t.Errorf("error write the response body: %v", err)
			}
		}
	}))

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		destineZipCode := r.URL.Query().Get("cepDestino")
		if len(destineZipCode) < 8 {
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte(`{"message":"invalid destine zip code"}`))
			if err != nil {
				t.Errorf("error write the response body: %v", err)
			}
		} else {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write(buildDeliveryDueDateBytes())
			if err != nil {
				t.Errorf("error write the response body: %v", err)
			}
		}
	}))
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
