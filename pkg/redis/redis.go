package redis

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/go-redis/redis/v8"
	log "github.com/gookit/slog"
	"github.com/koyote/pkg/config"
	"github.com/koyote/pkg/telegram"
)

var ctx = context.Background()
var redisClient *redis.Client

func ConnectToRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.GlobalAppConfig.Redis.Host, config.GlobalAppConfig.Redis.Port),
		Password: "",
		DB:       0,
	})

	err := redisClient.Set(ctx, "check", "connection", 5*time.Second).Err()
	if err != nil {
		log.Fatal("Error while connect to redis. Error: ", err)
	}

	sub := redisClient.Subscribe(ctx, "events")
	defer sub.Close()

	log.Info("Redis connection established! Subscribed for EVENTS channel!")
	ch := sub.Channel()

	for msg := range ch {
		log.Info("Received Event message from Redis. Trying to resend it to Telegram")
		ResendMessageToTelegram(msg)
	}

}

func PublishEventToRedisChannel(message string) {
	log.Info("Saving received event to Redis")
	redisClient.Publish(ctx, "events", message)
}

func ResendMessageToTelegram(msg *redis.Message) {
	eventMessage := strings.SplitAfter(msg.Payload, "|")
	chatID := strings.ReplaceAll(strings.ReplaceAll(eventMessage[0], "|", ""), "chatID:", "")
	message := strings.Replace(eventMessage[1], "message:", "", 1)

	err := retry.Do(
		func() error {
			err := telegram.SendEventMessage(chatID, message)
			return err
		})
	if err != nil {
		log.Error("Resending message failed. Event probably lost!")
	}

}
