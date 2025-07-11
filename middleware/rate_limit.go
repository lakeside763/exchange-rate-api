package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lakeside763/exchange-rate-api/storage"
)

type RateLimitMiddleware struct {
	Storage storage.Storage
}

func NewRateLimitMiddleware(store storage.Storage) *RateLimitMiddleware {
	return &RateLimitMiddleware{
		Storage: store,
	}
}

func (rl *RateLimitMiddleware) RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-Key")
			if apiKey == "" {
				http.Error(w, "Missing API Key", http.StatusBadRequest)
				return
			}

			user, ok := rl.Storage.GetUserByAPIKey(apiKey)
			if !ok {
				http.Error(w, "Invalid API Key", http.StatusUnauthorized)
				return
			}

			plan, ok := rl.Storage.GetPlanByID(user.PlanID)
			if !ok {
				http.Error(w, "Subscription plan not found", http.StatusInternalServerError)
				return
			}

			now := time.Now()
			dailyKey := fmt.Sprintf("ratelimit:%s:daily:%s", user.ID, now.Format("20060102"))
			weeklyKey := fmt.Sprintf("ratelimit:%s:weekly:%dW%d", user.ID, now.Year(), getWeek(now))
			monthlyKey := fmt.Sprintf("ratelimit:%s:monthly:%s", user.ID, now.Format("200601"))

			if !checkLimit(dailyKey, plan.DailyLimit, 24*time.Hour) ||
					!checkLimit(weeklyKey, plan.WeeklyLimit, 7*24*time.Hour) ||
					!checkLimit(monthlyKey, plan.MonthlyLimit, 31*24*time.Hour) {
						return
			}

			next.ServeHTTP(w, r)
		},
	)
}

func getWeek(t time.Time) int {
	_, week := t.ISOWeek()
	return week
}

// key
// limit
// ttl
func checkLimit(key string, limit int, ttl time.Duration) bool {
	count, err := storage.RedisClient.Incr(storage.Ctx(), key).Result()
	if err != nil {
		return false
	}

	if count == 1 {
		storage.RedisClient.Expire(storage.Ctx(), key, ttl)
	}

	return int(count) <= limit
}

func checkLimit2(key string, limit int) bool {
	storage.RateLimitMu.Lock()
	defer storage.RateLimitMu.Unlock()

	count := storage.RequestCounts[key]
	if count > limit {
		return false
	}

	storage.RequestCounts[key] = count + 1
	return true
}




// ratelimit:{user_id}:daily:20250707
// ratelimit:{user_id}:weekly:2025W28
// ratelimit:{user_id}:monthly:202507