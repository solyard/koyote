package redis

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/go-redis/redis/v8"
	log "github.com/gookit/slog"
	"github.com/jasonlvhit/gocron"
	"github.com/koyote/pkg/config"
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

	cronInterval := config.GlobalAppConfig.Redis.CheckUnsendedEventsInterval

	gocron.NewJob(uint64(cronInterval) * uint64(time.Second)).Do(EventScheduleCheckAndResendToTelegram)
}

func SaveEventMessageToCache(chatID int64, message string) {
	randomKey, err := rand.Int(rand.Reader, big.NewInt(100))
	cacheTTL := config.GlobalAppConfig.Redis.UnsendendTaskTTL
	redisClient.Set(ctx, fmt.Sprintf("%v-%v", string(chatID), randomKey), message, time.Duration(cacheTTL)*time.Second)
	if err != nil {
		log.Error("Error while saving task into redis. Message can be lost. Error: ", err)
		return
	}
}

func EventScheduleCheckAndResendToTelegram(chatID int64) {
	log.Info("STUB for event task iteration from redis")
	return
}
