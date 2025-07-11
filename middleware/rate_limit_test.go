package middleware

import (
	"testing"
	"time"

	"github.com/lakeside763/exchange-rate-api/storage"
	"github.com/stretchr/testify/assert"
)

func TestCheckLimit_WithRedis(t *testing.T) {
	storage.InitRedis()

	key := "test:ratelimit:key"
	limit := 3
	ttl := 5 * time.Second

	// Clean up
	storage.RedisClient.Del(storage.Ctx(), key)

	for i := 1; i <= limit; i++ {
		assert.True(t, checkLimit(key, limit, ttl), "Expected checkLimit to return true on attempt %d", i)
	}
	assert.False(t, checkLimit(key, limit, ttl), "Expected checkLimit to return false after reaching limit")
}
