package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type (
	RedisService interface {
		Get(key string) (string, error)
		GetTTL(key string) (*time.Duration, error)
		Set(key string, value interface{}) error
		SetWithExpire(key string, value interface{}) error
	}

	RedisImp struct {
		client *redis.Client
	}
)

func NewRedisConnection() (*redis.Client, error) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "192.168.56.5:6379",
		DB:   0,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("error cause: %+v\n", err)
		return nil, err
	}
	var statusCmd = rdb.ConfigSet(ctx, "notify-keyspace-events", "KEAx")
	log.Printf("status set expire:%+v\n", statusCmd)
	return rdb, nil
}

func NewClient(rdb *redis.Client) RedisService {
	return &RedisImp{
		client: rdb,
	}
}

func (service *RedisImp) Get(key string) (string, error) {
	ctx := context.Background()
	result, err := service.client.Get(ctx, key).Result()
	if err == redis.Nil {
		log.Println("key doesn't exist")
		return "", err
	} else if err != nil {
		log.Printf("error cause: %+v\n", err)
		return "", err
	}
	fmt.Printf("getting data with result:%s\n", result)
	return result, nil
}

func (service *RedisImp) GetTTL(key string) (*time.Duration, error) {
	ctx := context.Background()
	result, err := service.client.TTL(ctx, key).Result()
	if err == redis.Nil {
		log.Println("key doesn't exist")
		return nil, err
	} else if err != nil {
		log.Printf("error cause: %+v\n", err)
		return nil, err
	}
	fmt.Printf("getting data with result:%+v\n", result)
	return &result, nil
}

func (service *RedisImp) Set(key string, value interface{}) error {
	ctx := context.Background()
	result, err := service.client.Set(ctx, key, value, 0).Result()
	if err != nil {
		log.Printf("error cause: %+v\n", err)
		return err
	}
	fmt.Printf("set data with result:%s\n", result)
	return nil
}

func (service *RedisImp) SetWithExpire(key string, value interface{}) error {
	ctx := context.Background()
	result, err := service.client.SetEX(ctx, key, value, time.Second*1800).Result()
	if err != nil {
		log.Printf("error cause: %+v\n", err)
		return err
	}
	fmt.Printf("set data with result:%v\n", result)
	return nil
}
