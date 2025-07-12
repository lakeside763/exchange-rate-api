package storage

func (s *StorageService) GetExchangeRate(currencyPair string) (float64, bool) {
	rate, exists := s.ExchangeRates[currencyPair]
	if !exists {
		return 0, false
	}
	return rate, true
}
