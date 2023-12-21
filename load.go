package redis_cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func (c *Cache) LoadOneValue(id int, trigger string) (string, bool) {
	if trigger != "" {
		result, err := c.RedisClient.HGet(context.Background(), trigger, strconv.Itoa(id)).Result()
		if err == redis.Nil {
			return "", false
		} else if err != nil {
			return "", false
		} else {
			return result, true
		}
	}
	return "", false
}

func (c *Cache) LoadAllValues(trigger string) (map[string]string, bool) {
	if trigger != "" {
		result, err := c.RedisClient.HGetAll(context.Background(), trigger).Result()
		if err == redis.Nil {
			return nil, false
		} else if err != nil {
			return nil, false
		} else {
			return result, true
		}
	}
	return nil, false
}
