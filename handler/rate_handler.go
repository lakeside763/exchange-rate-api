package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ExchangeRateHandler struct {
	ExchangeRates map[string]float64
}

var exchangeRates = map[string]float64{
	"USD_EUR": 0.9,
	"EUR_USD": 1.1,
	"USD_GBP": 0.78,
	"GBP_USD": 1.28,
	"EUR_GBP": 0.86,
	"GBP_EUR": 1.16,
}

func NewExchangeRateHandler() *ExchangeRateHandler {
	return &ExchangeRateHandler{
		ExchangeRates: exchangeRates,
	}
}

func (h *ExchangeRateHandler) GetExchangeRate(w http.ResponseWriter, r *http.Request) {
	base := r.URL.Query().Get("base")
	target := r.URL.Query().Get("target")

	if base == "" || target == "" {
		http.Error(w, "Missing base or target currency", http.StatusBadRequest)
		return
	}

	key := fmt.Sprintf("%s_%s", base, target)
	rate, exists := h.ExchangeRates[key]
	if !exists {
		http.Error(w, "Exchange rate not found", http.StatusNotFound)
		return
	}

	response  := map[string]interface{}{
		"base": base,
		"target": target,
		"rate": rate,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}