package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lakeside763/exchange-rate-api/storage"
	"github.com/stretchr/testify/assert"
)

func setup() {
	storage.InitRedis()
}

func TestGetExchageRate_Success(t *testing.T) {
	setup()

	req, _ := http.NewRequest("GET", "/api/v1/rates?base=USD&target=EUR", nil)
	req.Header.Set("X-API-Key", "test123")

	h := NewExchangeRateHandler()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetExchangeRate)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), `"rate":0.9`)
}

func TestGetExchangeRate_MissingBase(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/rates?target=EUR", nil)
	req.Header.Set("X-API-Key", "test123")

	h := NewExchangeRateHandler()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetExchangeRate)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "Missing base or target")
}

func TestExchangeRate_InvalidCurrencyPair(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/rates?base=ABC&target=XYZ", nil)
	req.Header.Set("X-API-Key", "test123")

	h := NewExchangeRateHandler()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetExchangeRate)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Contains(t, rr.Body.String(), "Exchange rate not found")
}
