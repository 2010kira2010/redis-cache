package redis_cache

import (
	"context"
	"encoding/json"
)

func (c *Cache) StoreValue(value interface{}, trigger string, id int) {
	loadValue, errLoadValue := c.LoadOneValue(id, trigger)
	if errLoadValue {
		dataTwo, _ := json.Marshal(value)
		mergedJSON, err := mergeInterface(loadValue, string(dataTwo))
		if err == nil {
			c.RedisClient.HSet(context.Background(), trigger, int64(id), mergedJSON)
		}
	} else {
		data, err := json.Marshal(value)
		if err == nil {
			c.RedisClient.HSet(context.Background(), trigger, int64(id), string(data))
		}
	}
}
