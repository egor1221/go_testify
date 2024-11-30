package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
    req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)

    // здесь нужно добавить необходимые проверки
	body := responseRecorder.Body.String()
    list := strings.Split(body, ",")

	assert.NotEmpty(t, body)
	assert.Len(t, list, totalCount)
}

func TestMainHandlerWhenAnotherCity(t *testing.T) {
    actualCity := "moscow"
    req := httptest.NewRequest("GET", "/cafe?count=4&city=brest", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)

    // здесь нужно добавить необходимые проверки
    city := req.URL.Query().Get("city")

	assert.Equal(t, city, actualCity)
}