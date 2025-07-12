package main

import (
	"log"
	"net/http"

	"github.com/lakeside763/exchange-rate-api/handler"
	"github.com/lakeside763/exchange-rate-api/middleware"
	"github.com/lakeside763/exchange-rate-api/storage"
)

func main() {
	storage.InitRedis()

	storage := storage.NewStorageService()
	handler := handler.NewExchangeRateHandler(storage)
	middleware := middleware.NewRateLimitMiddleware(storage)

	mux := http.NewServeMux()
	mux.Handle("/api/v1/rates", middleware.RateLimit(http.HandlerFunc(handler.GetExchangeRate)))

	log.Println("Starting server on :5200...")
	log.Fatal(http.ListenAndServe(":5200", mux))
}