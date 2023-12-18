# Go Package for Redis
<p>
    <a href="https://github.com/2010kira2010/redis-cache">
        <img src="https://img.shields.io/badge/codecov-55.1%25-success" alt="Code Coverage">
    </a>
    <a href="https://github.com/2010kira2010/redis-cache">
        <img src="https://img.shields.io/badge/build-passes-success" alt="Build Status">
    </a>
    <a href="https://github.com/2010kira2010/redis-cache/blob/main/LICENSE">
        <img src="https://img.shields.io/github/license/saltstack/salt" alt="License">
    </a>
</p> 
This package provides a Golang client-cache for Redis.

## Installation
`go get github.com/2010kira2010/redis-cache`

## Quick Start
```go
package main

import (
	"fmt"
	redis_cache "github.com/2010kira2010/redis-cache"
	"github.com/redis/go-redis/v9"
)

func NewCache() *redis_cache.Cache {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "1.1.1.1:6379", // address and port Redis-server
		Password: "",                 // password Redis-server, if there
		DB:       0,                  // Redis database number to be used
	})

	return &redis_cache.Cache{
		RedisClient: redisClient,
	}
}

var Cache *redis_cache.Cache = NewCache()

type Lead struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func main() {
	Cache.StoreValue(Lead{
		ID:       111111,
		Name:     "Test",
	}, "test", 111111)

	var leadOne Lead
	a, _ := Cache.LoadOneValue(&leadOne, 111111, "test")
	fmt.Println(a)
}

```