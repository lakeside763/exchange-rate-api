package storage

import "github.com/lakeside763/exchange-rate-api/model"

type StorageService struct {
	Users map[string]model.User
	Plans map[string]model.Plan
}

func NewStorageService() *StorageService {
	return &StorageService{
		Users: Users,
		Plans: Plans,
	}
}
 
func (s *StorageService) GetUserByAPIKey(apiKey string) (*model.User, bool) {
	user, ok := s.Users[apiKey]
	if !ok {
		return nil, false
	}

	return &user, true
}


func (s *StorageService) GetPlanByID(planID string) (*model.Plan, bool) {
	plan, ok := s.Plans[planID]
	if !ok {
		return nil, false
	}

	return &plan, true
}