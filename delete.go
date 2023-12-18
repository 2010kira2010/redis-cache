package redis_cache

import (
	"context"
	"strconv"
)

func (c *Cache) DeleteOneValue(trigger string, id int) {
	c.RedisClient.HDel(context.Background(), trigger, strconv.Itoa(id))
}

func (c *Cache) DeleteAllValues(trigger string) {
	c.RedisClient.Del(context.Background(), trigger)
}
