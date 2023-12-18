package redis_cache

import (
	"context"
	"encoding/json"
)

func (c *Cache) StoreValue(value interface{}, trigger string, id int) {
	loadValue, errLoadValue := c.LoadOneValue(value, id, trigger)
	if errLoadValue {
		dataOne, _ := json.Marshal(loadValue)
		dataTwo, _ := json.Marshal(value)
		mergedJSON, err := mergeInterface(string(dataOne), string(dataTwo))
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
