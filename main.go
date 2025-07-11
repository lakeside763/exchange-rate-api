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

	handler := handler.NewExchangeRateHandler()
	service := storage.NewStorageService()
	middleware := middleware.NewRateLimitMiddleware(service)

	mux := http.NewServeMux()
	mux.Handle("/api/v1/rates", middleware.RateLimit(http.HandlerFunc(handler.GetExchangeRate)))

	log.Println("Starting server on :5200...")
	log.Fatal(http.ListenAndServe(":5200", mux))
}