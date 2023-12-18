package redis_cache

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func (c *Cache) LoadOneValue(value interface{}, id int, trigger string) (interface{}, bool) {
	if trigger != "" {
		result, err := c.RedisClient.HGet(context.Background(), trigger, strconv.Itoa(id)).Result()
		if err == redis.Nil {
			return nil, false
		} else if err != nil {
			return nil, false
		} else {
			err := json.Unmarshal([]byte(result), value)
			if err == nil {
				return value, true
			}
		}
	}
	return nil, false
}

func (c *Cache) LoadAllValues(value interface{}, trigger string) (map[int]interface{}, bool) {
	values := make(map[int]interface{})
	if trigger != "" {
		results, err := c.RedisClient.HGetAll(context.Background(), trigger).Result()
		if err == nil {
			for i, result := range results {
				err := json.Unmarshal([]byte(result), &value)
				if err == nil {
					i, _ := strconv.Atoi(i)
					values[i] = value
				}
			}
		} else {
			return nil, false
		}
	}
	if len(values) != 0 {
		return values, true
	}
	return nil, false
}
