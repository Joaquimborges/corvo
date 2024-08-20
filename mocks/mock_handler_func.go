package mocks

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type requestValidator func(*http.Request) bool

func BuildTestHandleFunc(
	shouldReturnError requestValidator,
	errorResponse []byte,
	successResponse []byte,
	successStatusCode int,
	errorStatusCode int,
	t *testing.T,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if shouldReturnError(r) {
			w.WriteHeader(errorStatusCode)
			_, err := w.Write(errorResponse)
			if err != nil {
				t.Errorf("error write the response body: %v", err)
			}
		} else {
			w.WriteHeader(successStatusCode)
			_, err := w.Write(successResponse)
			if err != nil {
				t.Errorf("error write the response body: %v", err)
			}
		}
	}
}

func BuildGenerateAccessTokenTestSever(t *testing.T) *httptest.Server {
	returnErrorIf := func(req *http.Request) bool {
		return req.Header.Get("authorization") != "Basic foo"
	}

	errorResponseBody := []byte(`{"message":"bad credential"}`)
	successResponseBody := []byte(`{"ambiente":"sandbox", "cnpj":"67412064000170", "token":"bc1e0e7d-60f5-4cad-a30c-480a64405d27"}`)
	return httptest.NewServer(BuildTestHandleFunc(
		returnErrorIf,
		errorResponseBody,
		successResponseBody,
		http.StatusCreated,
		http.StatusUnauthorized,
		t,
	))
}
