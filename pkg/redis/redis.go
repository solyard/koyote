package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	log "github.com/gookit/slog"
)

var ctx = context.Background()
var redisClient *redis.Client

func ConnectToRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := redisClient.Set(ctx, "check", "connection", 5*time.Second).Err()
	if err != nil {
		log.Fatal("Error while connect to redis. Error: ", err)
	}

	log.Info("Redis connection established!")
}

func SaveEventMessageToCache(chatID int64, message string) {
	err := redisClient.Set(ctx, string(chatID), message, 60*time.Minute)
	if err != nil {
		log.Error("Error while saving task into redis. Message can be lost. Error: ", err)
		return
	}
}
