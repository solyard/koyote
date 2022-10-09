package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	log "github.com/gookit/slog"
)

var ctx = context.Background()

func ConnectToRedis() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := redisClient.Set(ctx, "check", "connection", 5*time.Second).Err()
	if err != nil {
		log.Error("Error while connect to redis. Error: ", err)
	}
}

func AddTaskToPool() {
	return
}

func RemoveTaskFromPool() {
	return
}
