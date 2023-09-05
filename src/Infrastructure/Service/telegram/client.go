package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	Bot *tgbotapi.BotAPI
}

func CreateBot() *TelegramBot {
	botToken := "6644558420:AAHO2q9oYRbv_8Jh6xLt9oTK0SMLoU9lRnU"

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	config := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     "/list",
			Description: "All teams",
		},
		tgbotapi.BotCommand{
			Command:     "/myTeams",
			Description: "Subscribed teams",
		},
		tgbotapi.BotCommand{
			Command:     "/subscribe",
			Description: "Subscribe",
		},
		tgbotapi.BotCommand{
			Command:     "/unsubscribe",
			Description: "Unsubscribe",
		},
	)

	bot.Request(config)

	return &TelegramBot{Bot: bot}
}

func (bot *TelegramBot) ReadChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return bot.Bot.GetUpdatesChan(u)
}
