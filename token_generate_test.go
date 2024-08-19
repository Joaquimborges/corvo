package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateAccessToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	defer server.Close()

	t.Run("should return error, invalid authorization key", func(t *testing.T) {
		config := &Config{
			PostCard:          "00112233",
			AuthorizationCode: "bar",
			UrlMapper: map[urlKey]string{
				GenerateAccessTokenUrlKey: server.URL,
			},
		}
		tokenData, err := generateAccessToken(config, NewHttpClient())

		require.Error(t, err)
		require.Contains(t, err.Error(), "bad credential")
		require.Nil(t, tokenData)
	})

	t.Run("shold return ok", func(t *testing.T) {
		config := &Config{
			PostCard:          "00112233",
			AuthorizationCode: "foo",
			UrlMapper: map[urlKey]string{
				GenerateAccessTokenUrlKey: server.URL,
			},
		}
		tokenData, err := generateAccessToken(config, NewHttpClient())

		expectedTokenData := buildTokenData()

		require.NoError(t, err)
		require.NotNil(t, tokenData)
		require.Equal(t, expectedTokenData.Cnpj, tokenData.Cnpj)
		require.Equal(t, expectedTokenData.Environment, tokenData.Environment)
		require.Equal(t, config.PostCard, tokenData.PostCard.Number)
		require.Equal(t, expectedTokenData.Token, tokenData.Token)
	})
}

func buildTokenResponseBytes(t *testing.T, data *tokenData) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		t.Errorf("error marshal token data: %v", err)
	}
	return bytes
}

func buildTokenData() *tokenData {
	return &tokenData{
		Environment: "sandbox",
		Cnpj:        "67.412.064/0001-70",
		PostCard: postCard{
			Number: "00112233",
		},
		Token: "bc1e0e7d-60f5-4cad-a30c-480a64405d27",
	}
}
