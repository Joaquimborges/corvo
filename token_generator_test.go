package main

import (
	"testing"

	"github.com/joaquimborges/corvo/mocks"
)

func TestTokenGenerator(t *testing.T) {
	t.Run("should return error, invalid config", func(t *testing.T) {
		config := &Config{
			PostCard:          "foo",
			AuthorizationCode: "bar",
		}

		client := mocks.NewClientMock(true)
		rest := NewHttpClient()
		token, err := generateAccessToken(config)
	})
}
