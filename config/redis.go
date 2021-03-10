package config

import "github.com/go-redis/redis"

type RedisDB struct {
	DBLive redis.Client
}

func NewRedisDatabase() *RedisDB {
	client := redis.NewClient(&redis.Options{
		Addr:     REDISDBURL,
		Password: "",
		DB:       0,
	})
	return &RedisDB{
		DBLive: *client,
	}
}
