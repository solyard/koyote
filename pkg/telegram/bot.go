package telegram

import (
	log "github.com/gookit/slog"
	"github.com/koyote/pkg/config"
	"github.com/koyote/pkg/redis"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

var Bot *telego.Bot

func StartBot() {
	// Initialise BOT
	botToken := config.GlobalAppConfig.Global.TelegramBotToken
	bot, err := telego.NewBot(botToken)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Telegram bot started!")
	Bot = bot
	// Start Polling
	_, _ = bot.UpdatesViaLongPulling(nil)
	defer bot.StopLongPulling()
}

func SendEventMessage(chatID int64, eventMessage string) {
	log.Info("Received event message. Message: ", eventMessage)
	_, err := Bot.SendMessage(
		&telego.SendMessageParams{
			ChatID:                tu.ID(chatID),
			Text:                  eventMessage,
			ParseMode:             "HTML",
			DisableWebPagePreview: true,
		},
	)

	if err != nil && config.GlobalAppConfig.Redis.Enabled {
		log.Error("Error while sending message to telegram. Save task in Redis. Error: ", err)
		redis.SaveEventMessageToCache(chatID, eventMessage)
		return
	} else if err != nil {
		log.Error("Error while sending message to telegram. Redis disabled so you missed this message in telegram. Error: ", err)
		return
	}
}
