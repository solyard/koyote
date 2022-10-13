package telegram

import (
	log "github.com/gookit/slog"
	"github.com/koyote/pkg/config"
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

func SendEventMessage(chatID int64, eventMessage string) error {
	log.Info("Received event message. Message: ", eventMessage)
	_, err := Bot.SendMessage(
		&telego.SendMessageParams{
			ChatID:                tu.ID(chatID),
			Text:                  eventMessage,
			ParseMode:             "HTML",
			DisableWebPagePreview: true,
		},
	)

	if err != nil {
		return err
	} else {
		return nil
	}
}
