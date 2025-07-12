package storage

import (
	"github.com/lakeside763/exchange-rate-api/model"
)

var Plans = map[string]model.Plan{
	"free": {
		ID:           "free",
		Name:         "Free Plan",
		DailyLimit:   100,
		WeeklyLimit:  500,
		MonthlyLimit: 2000,
		PriceUSD:     0.0,
	},
	"pro": {
		ID:           "pro",
		Name:         "Pro Plan",
		DailyLimit:   1000,
		WeeklyLimit:  5000,
		MonthlyLimit: 20000,
		PriceUSD:     49.99,
	},
}

var Users = map[string]model.User{
	"test123": {
		ID:     "u1",
		Email:  "demo@example.com",
		APIKey: "test123",
		PlanID: "free",
	},
	"pro456": {
		ID:     "u2",
		Email:  "pro@example.com",
		APIKey: "pro456",
		PlanID: "pro",
	},
}

var ExchangeRates = map[string]float64{
	"USD_EUR": 0.9,
	"EUR_USD": 1.1,
	"USD_GBP": 0.78,
	"GBP_USD": 1.28,
	"EUR_GBP": 0.86,
	"GBP_EUR": 1.16,
}

// var (
// 	RequestCounts = make(map[string]int)
// 	RateLimitMu            sync.Mutex
// )
