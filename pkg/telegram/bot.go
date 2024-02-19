package telegram

import (
	"fmt"
	"strconv"

	log "github.com/gookit/slog"
	"github.com/koyote/pkg/config"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/pkg/errors"
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
	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()

	// Create bot handler and specify from where to get updates
	bh, _ := th.NewBotHandler(bot, updates)

	// Stop handling updates
	defer bh.Stop()

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _ = bot.SendMessage(
			&telego.SendMessageParams{
				ChatID:             update.Message.Chat.ChatID(),
				Text:               fmt.Sprintf("Hello %s! Use /help command to get available commands!", update.Message.From.FirstName),
				ParseMode:          "HTML",
				LinkPreviewOptions: &telego.LinkPreviewOptions{IsDisabled: true},
				ReplyParameters:    &telego.ReplyParameters{MessageID: update.Message.MessageID},
			},
		)
	}, th.CommandEqual("start"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _ = bot.SendMessage(
			&telego.SendMessageParams{
				ChatID: update.Message.Chat.ChatID(),
				Text: fmt.Sprintf(`
				Hello %s! Have a look on supported commands:
	
	/chatID - Return current chatID
	/threadID - Return current threadID (If 0 then you are in General Thread or chat without Threads support)`, update.Message.From.FirstName),
				ParseMode:          "HTML",
				LinkPreviewOptions: &telego.LinkPreviewOptions{IsDisabled: true},
				ReplyParameters:    &telego.ReplyParameters{MessageID: update.Message.MessageID},
			},
		)
	}, th.CommandEqual("help"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _ = bot.SendMessage(
			&telego.SendMessageParams{
				ChatID:             update.Message.Chat.ChatID(),
				Text:               fmt.Sprintf("Current Thread ID: <code>%v</code>", update.Message.MessageThreadID),
				ParseMode:          "HTML",
				LinkPreviewOptions: &telego.LinkPreviewOptions{IsDisabled: true},
				ReplyParameters:    &telego.ReplyParameters{MessageID: update.Message.MessageID},
			},
		)
	}, th.CommandEqual("threadID"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _ = bot.SendMessage(
			&telego.SendMessageParams{
				ChatID:             update.Message.Chat.ChatID(),
				Text:               fmt.Sprintf("Current chat ID: <code>%v</code>", update.Message.Chat.ID),
				ParseMode:          "HTML",
				LinkPreviewOptions: &telego.LinkPreviewOptions{IsDisabled: true},
				ReplyParameters:    &telego.ReplyParameters{MessageID: update.Message.MessageID},
			},
		)
	}, th.CommandEqual("chatID"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _ = bot.SendMessage(
			&telego.SendMessageParams{
				ChatID:             update.Message.Chat.ChatID(),
				Text:               "Unknown command, use /help to get additional info about commands",
				ParseMode:          "HTML",
				LinkPreviewOptions: &telego.LinkPreviewOptions{IsDisabled: true},
				ReplyParameters:    &telego.ReplyParameters{MessageID: update.Message.MessageID},
			},
		)
	}, th.AnyCommand())

	// Start handling updates
	bh.Start()
}

func SendEventMessage(chatID string, eventMessage string) error {
	chatIDInt, err := strconv.Atoi(chatID)
	if err != nil {
		return err
	}
	_, err = Bot.SendMessage(
		&telego.SendMessageParams{
			ChatID:             telego.ChatID{ID: int64(chatIDInt)},
			Text:               eventMessage,
			ParseMode:          "HTML",
			LinkPreviewOptions: &telego.LinkPreviewOptions{IsDisabled: true},
		},
	)

	return errors.Wrap(err, "Error while sending message to Telegram")
}

func SendEventMessageToThread(chatID, threadID, eventMessage string) error {
	chatIDInt, err := strconv.Atoi(chatID)
	if err != nil {
		return err
	}

	threadIDInt, err := strconv.Atoi(threadID)
	if err != nil {
		return err
	}

	_, err = Bot.SendMessage(
		&telego.SendMessageParams{
			ChatID:             telego.ChatID{ID: int64(chatIDInt)},
			MessageThreadID:    int(threadIDInt),
			Text:               eventMessage,
			ParseMode:          "HTML",
			LinkPreviewOptions: &telego.LinkPreviewOptions{IsDisabled: true},
		},
	)

	return errors.Wrap(err, "Error while sending message to Telegram")
}
