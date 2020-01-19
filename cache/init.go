package cache

import (
	"github.com/go-redis/redis"
)

func InitRedis(addr,pass string,db int) (*redis.Client,error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})

	_, err := client.Ping().Result()

	return client,err
}
