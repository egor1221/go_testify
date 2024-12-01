package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


func TestMainHandlerWhenOK(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=20&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    require.Equal(t, responseRecorder.Code, http.StatusOK)
}
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
    req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)


	body := responseRecorder.Body.String()
    list := strings.Split(body, ",")

    if len(list) < totalCount {
        totalCount = len(list)
    }

	assert.NotEmpty(t, body)
	assert.Equal(t, len(list), totalCount)
}

func TestMainHandlerWhenAnotherCity(t *testing.T) {
    actualCity := "moscow"
    req := httptest.NewRequest("GET", "/cafe?count=4&city=brest", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    city := req.URL.Query().Get("city")

	assert.Equal(t, city, actualCity)
}