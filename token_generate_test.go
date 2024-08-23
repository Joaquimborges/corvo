package corvo

import (
	"testing"

	"github.com/Joaquimborges/corvo/mocks"
	"github.com/stretchr/testify/require"
)

func TestGenerateAccessToken(t *testing.T) {
	server := mocks.BuildGenerateAccessTokenTestSever(t)
	defer server.Close()

	t.Run("should return error, invalid authorization key", func(t *testing.T) {
		config := &Config{
			PostCard:          "00112233",
			AuthorizationCode: "bar",
			UrlMapper: map[EndpointURL]string{
				GenerateAccessTokenURL: server.URL,
			},
		}
		tokenData, err := generateAccessToken(config, newHttpClient())

		require.Error(t, err)
		require.Contains(t, err.Error(), "bad credential")
		require.Nil(t, tokenData)
	})

	t.Run("should return ok", func(t *testing.T) {
		config := &Config{
			PostCard:          "00112233",
			AuthorizationCode: "foo",
			UrlMapper: map[EndpointURL]string{
				GenerateAccessTokenURL: server.URL,
			},
		}
		expectedTokenData := buildTokenData()

		tokenData, err := generateAccessToken(config, newHttpClient())

		require.NoError(t, err)
		require.NotNil(t, tokenData)
		require.Equal(t, expectedTokenData.Cnpj, tokenData.Cnpj)
		require.Equal(t, expectedTokenData.Environment, tokenData.Environment)
		require.Equal(t, expectedTokenData.Token, tokenData.Token)
	})
}

func buildTokenData() *tokenData {
	return &tokenData{
		Environment: "sandbox",
		Cnpj:        "67412064000170",
		Token:       "bc1e0e7d-60f5-4cad-a30c-480a64405d27",
	}
}
