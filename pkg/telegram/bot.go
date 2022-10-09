package telegram

import (
	"os"

	log "github.com/gookit/slog"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

var Bot *telego.Bot

func StartBot() {
	// Initialise BOT
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	bot, err := telego.NewBot(botToken)
	if err != nil {
		log.Fatal(err)
	}

	// Start Polling
	_, _ = bot.UpdatesViaLongPulling(nil)
	defer bot.StopLongPulling()
}

func sendEventMessage(chatID int64, eventMessage string) {
	Bot.SendMessage(
		tu.Message(
			tu.ID(chatID),
			eventMessage,
		),
	)
}
