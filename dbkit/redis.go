package dbkit

import (
	"log"

	"github.com/go-redis/redis"
)

func NewRedisClient(opt *redis.Options) *redis.Client {
	redisdb := redis.NewClient(opt)

	log.Println("redis status:", redisdb.Ping())

	return redisdb
}
