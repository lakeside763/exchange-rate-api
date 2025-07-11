package storage

import "github.com/lakeside763/exchange-rate-api/model"

type Storage interface {
	GetUserByAPIKey(apiKey string) (*model.User, bool)
	GetPlanByID(planID string) (*model.Plan, bool)
}