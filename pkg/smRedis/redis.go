package smRedis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var (
	Rdb *redis.Client
	Ctx = context.Background()
)

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // если без пароля
		DB:       0,
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("could not connect to Redis: %v", err)
	}

	log.Println("Connected to Redis")
}
