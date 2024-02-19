package redis_cache

import (
	"context"
	"strconv"
)

func (c *Cache) DeleteOneValue(trigger, id interface{}) {
	c.RedisClient.HDel(context.Background(), trigger.(string), strconv.FormatFloat(id.(float64), 'f', -1, 64))
}

func (c *Cache) DeleteAllValues(trigger interface{}) {
	c.RedisClient.Del(context.Background(), trigger.(string))
}
