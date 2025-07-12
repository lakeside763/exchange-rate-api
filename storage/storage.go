package storage

import "github.com/lakeside763/exchange-rate-api/model"

type Storage interface {
	GetUserByAPIKey(apiKey string) (*model.User, bool)
	GetPlanByID(planID string) (*model.Plan, bool)
	GetExchangeRate(currencyPair string) (float64, bool)
}

type StorageService struct {
	Users map[string]model.User
	Plans map[string]model.Plan
	ExchangeRates map[string]float64
}

func NewStorageService() *StorageService {
	return &StorageService{
		Users: Users,
		Plans: Plans,
		ExchangeRates: ExchangeRates,
	}
}