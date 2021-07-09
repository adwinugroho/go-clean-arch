package subscribe

import (
	"context"
	redisHelper "go-clean-arch/pkg/redis"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func ValidateExpireUID() {
	var ctx = context.Background()
	getConnectRedis, err := redisHelper.NewRedisConnection()
	if err != nil {
		log.Println("error connect to redis", err)
		return
	}
	pong, err := getConnectRedis.Ping(context.Background()).Result()
	if err != nil {
		log.Println(err)
	}

	log.Println(pong)

	var status = getConnectRedis.SetEX(context.Background(), "test", "value", 1*time.Minute)
	log.Println(status)

	pubSub := getConnectRedis.PSubscribe(ctx, "__keyevent@0__:expired")
	log.Printf("pubSub:%+v\n", pubSub)
	go func(*redis.PubSub) {
		for {
			msgi, err := pubSub.Receive(ctx)
			if err != nil {
				log.Println("error get receive", err)
				return
			}
			switch msg := msgi.(type) {
			case *redis.Message:
				// err := permission.PermsDeletepermdoc(msg.Payload)
				// if err != nil {
				// 	log.Println("error get receive", err)
				// 	return
				// }
				log.Printf("Message: %s %s\n", msg.Channel, msg.Payload) //msg.Payload == uid
			case *redis.Subscription:
				log.Printf("Subscription: %s %s %d\n", msg.Kind, msg.Channel, msg.Count)
				if msg.Count == 0 {
					log.Println("error count subs 0")
					return
				}
			case error:
				log.Printf("error, cause: %v\n", msg)
				return
			}
		}
	}(pubSub)

}
