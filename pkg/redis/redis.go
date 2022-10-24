package redis

import (
	"context"
	"fmt"
	"strings"
	"time"

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
		_, err = config.RedisCB.Execute(func() (interface{}, error) {
			result, err := ResendMessageToTelegram(msg)
			if err != nil {
				return result, err
			}
			return result, nil
		})
	}

}

func PublishEventToRedisChannel(message string) {
	redisClient.Publish(ctx, "events", message)
}

func ResendMessageToTelegram(msg *redis.Message) (bool, error) {
	eventMessage := strings.SplitAfter(msg.Payload, "|")
	chatID := strings.ReplaceAll(strings.ReplaceAll(eventMessage[0], "|", ""), "chatID:", "")
	message := strings.Replace(eventMessage[1], "message:", "", 1)
	err := telegram.SendEventMessage(chatID, message)
	if err != nil {
		return false, err
	}
	return true, nil
}
