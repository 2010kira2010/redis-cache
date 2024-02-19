package redis_cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func (c *Cache) LoadOneValue(id, trigger interface{}) (string, bool) {
	if trigger != "" {
		result, err := c.RedisClient.HGet(context.Background(), trigger.(string), strconv.FormatFloat(id.(float64), 'f', -1, 64)).Result()
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

func (c *Cache) LoadAllValues(trigger interface{}) (map[string]string, bool) {
	if trigger != "" {
		result, err := c.RedisClient.HGetAll(context.Background(), trigger.(string)).Result()
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
