package redis_cache

import (
	"context"
	"encoding/json"
)

func (c *Cache) StoreValue(value, trigger, id interface{}) {
	loadValue, errLoadValue := c.LoadOneValue(id, trigger)
	if errLoadValue {
		dataTwo, _ := json.Marshal(value)
		mergedJSON, err := mergeInterface(loadValue, string(dataTwo))
		if err == nil {
			c.RedisClient.HSet(context.Background(), trigger.(string), id.(float64), mergedJSON)
		}
	} else {
		data, err := json.Marshal(value)
		if err == nil {
			c.RedisClient.HSet(context.Background(), trigger.(string), id.(float64), string(data))
		}
	}
}
