package redis_cache

import (
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	RedisClient *redis.Client
}
