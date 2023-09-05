package telegram

import (
	"fmt"
	serviceOp "go-football/src/Application/Service"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Actions() {
	bot := CreateBot()
	fmt.Println("Bot created")

	updates := bot.ReadChannel()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		fmt.Println(update.Message.Text)

		// Check if the user sent a command.
		if update.Message.IsCommand() {
			chatId := update.Message.Chat.ID

			if update.Message.Command() == "list" {
				getAllTeams()
			}
			if update.Message.Command() == "myTeams" {
				getMyTeams(bot, chatId)
			}
			if update.Message.Command() == "subscribe" {
				subscribeOnTeam(bot, updates, chatId)
			}
			// if update.Message.Command() == "unsubscribe" {
			// 	UnubscribeFromTeam(updates, chatId)
			// }
		}
	}
}

func subscribeOnTeam(bot *TelegramBot, updates tgbotapi.UpdatesChannel, chatId int64) {
	var teamId int64
	// Wait for the user to reply with the event name.
	eventNameUpdate := <-updates
	if eventNameUpdate.Message != nil {
		teamId, _ = strconv.ParseInt(eventNameUpdate.Message.Text, 10, 64)
	}

	serviceOp.SubscribeOnTeam(int32(chatId), int32(teamId))

	message(bot, chatId, "Subscribed")
}

func getMyTeams(bot *TelegramBot, chatId int64) {
	teams := serviceOp.GetMyTeams(int32(chatId))
	fmt.Println(teams)

	message(bot, chatId, "My teams")
}

func getAllTeams(bot *TelegramBot, chatId int64) {
	teams := serviceOp.GetTeams()
	fmt.Println(teams)

	message(bot, chatId, "All teams")
}

func message(bot *TelegramBot, chatId int64, data string) {
	bot.Bot.Send(tgbotapi.NewMessage(chatId, data))
}
