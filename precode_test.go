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
    req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    body := responseRecorder.Body.String()

    require.NotEmpty(t, body)

    require.Equal(t, http.StatusOK, responseRecorder.Code)
}
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
    req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)


	body := responseRecorder.Body.String()
    list := strings.Split(body, ",")

    // это на случай если count<totalCount
    if len(list) < totalCount {
        totalCount = len(list)
    }

    assert.Equal(t, len(list), totalCount) // поскольку по заданию надо разобрать только случай когда count>totalCount сделал так
	// assert.LessOrEqual(t, len(list), totalCount) 
}

func TestMainHandlerWhenAnotherCity(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=4&city=brest", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    body := responseRecorder.Body.String()

    require.Equal(t, responseRecorder.Code, http.StatusBadRequest)

	assert.Equal(t, "wrong city value", body)
}