package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lakeside763/exchange-rate-api/storage"
)

type ExchangeRateHandler struct {
	Storage storage.Storage
}

func NewExchangeRateHandler(storage storage.Storage) *ExchangeRateHandler {
	return &ExchangeRateHandler{
		Storage: storage,
	}
}

func (h *ExchangeRateHandler) GetExchangeRate(w http.ResponseWriter, r *http.Request) {
	base := r.URL.Query().Get("base")
	target := r.URL.Query().Get("target")

	if base == "" || target == "" {
		http.Error(w, "Missing base or target currency", http.StatusBadRequest)
		return
	}

	fmt.Println("Fetching exchange rate for:", base, "to", target)

	key := fmt.Sprintf("%s_%s", base, target)
	fmt.Println("Key for exchange rate lookup:", key)
	// Debug: print all available keys (if using a map)
	// fmt.Println("Available keys:", h.Storage.ListKeys()) // implement ListKeys() if needed
	rate, exists := h.Storage.GetExchangeRate(key)
	if !exists {
		fmt.Println("Available keys do not contain:", key)
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