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

			if update.Message.Command() == "start" {
				saveUser(update.Message.Chat.UserName, chatId)
			}
			if update.Message.Command() == "list" {
				getAllTeams(bot, chatId)
			}
			if update.Message.Command() == "myTeams" {
				getMyTeams(bot, chatId)
			}
			if update.Message.Command() == "subscribe" {
				subscribeOnTeam(bot, updates, chatId)
			}
			if update.Message.Command() == "unsubscribe" {
				unsubscribeFromTeam(bot, updates, chatId)
			}
		}
	}
}

func subscribeOnTeam(bot *TelegramBot, updates tgbotapi.UpdatesChannel, chatId int64) {
	userId, err := serviceOp.GetUser(int32(chatId))
	if nil != err {
		message(bot, chatId, "Wrong user")
		return
	}

	message(bot, chatId, "Give me an ID of the team you want subscribe to")

	var teamId int64
	// Wait for the user to reply with the event name.
	eventNameUpdate := <-updates
	if eventNameUpdate.Message != nil {
		teamId, _ = strconv.ParseInt(eventNameUpdate.Message.Text, 10, 64)
	}

	serviceOp.SubscribeOnTeam(userId.Id, int32(teamId))

	message(bot, chatId, "Subscribed")
}

func unsubscribeFromTeam(bot *TelegramBot, updates tgbotapi.UpdatesChannel, chatId int64) {
	userId, err := serviceOp.GetUser(int32(chatId))
	if nil != err {
		message(bot, chatId, "Wrong user")
		return
	}

	message(bot, chatId, "Give me an ID of the team you want to unsubscribe")

	var teamId int64
	// Wait for the user to reply with the event name.
	eventNameUpdate := <-updates
	if eventNameUpdate.Message != nil {
		teamId, _ = strconv.ParseInt(eventNameUpdate.Message.Text, 10, 64)
	}

	serviceOp.UnubscribeFromTeam(userId.Id, int32(teamId))

	message(bot, chatId, "Unsubscribed")
}

func getMyTeams(bot *TelegramBot, chatId int64) {
	userId, err := serviceOp.GetUser(int32(chatId))
	if nil != err {
		message(bot, chatId, "Wrong user")
		return
	}

	teams := serviceOp.GetMyTeams(userId.Id)
	if len(teams) == 0 {
		message(bot, chatId, "No subscribed teams")
		return
	}

	var list string
	for _, team := range teams {
		list += strconv.Itoa(int(team.Id)) + " - " + team.Name + "\n"
	}

	fmt.Println(list)

	message(bot, chatId, list)
}

func getAllTeams(bot *TelegramBot, chatId int64) {
	teams := serviceOp.GetTeams()
	var list string
	for _, team := range teams {
		list += strconv.Itoa(int(team.Id)) + " - " + team.Name + "\n"
	}

	fmt.Println(list)

	message(bot, chatId, list)
}

func message(bot *TelegramBot, chatId int64, data string) {
	bot.Bot.Send(tgbotapi.NewMessage(chatId, data))
}

func saveUser(username string, chatId int64) {
	serviceOp.CreateUser(username, int32(chatId))
}
