package telegram

import (
	"strconv"

	log "github.com/gookit/slog"
	"github.com/koyote/pkg/config"
	"github.com/mymmrac/telego"
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

func SendEventMessage(chatID string, eventMessage string) error {
	chatIDInt, err := strconv.Atoi(chatID)
	if err != nil {
		return err
	}
	log.Info("Received event message for Telegram")
	_, err = Bot.SendMessage(
		&telego.SendMessageParams{
			ChatID:                telego.ChatID{ID: int64(chatIDInt)},
			Text:                  eventMessage,
			ParseMode:             "HTML",
			DisableWebPagePreview: true,
		},
	)

	return err
}
