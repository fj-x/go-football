package telegram

import (
	"fmt"
	serviceOp "go-football/src/Application/Service"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramActions struct {
	userService         *serviceOp.UserService
	teamService         *serviceOp.TeamService
	subscriptionService *serviceOp.SubscriptionService
	notificationService *serviceOp.NotificationService
}

func New(
	userService *serviceOp.UserService,
	teamService *serviceOp.TeamService,
	subscriptionService *serviceOp.SubscriptionService,
	notificationService *serviceOp.NotificationService,
) *TelegramActions {
	return &TelegramActions{
		userService:         userService,
		teamService:         teamService,
		subscriptionService: subscriptionService,
		notificationService: notificationService,
	}
}

func (actions TelegramActions) Actions() {
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
				actions.saveUser(update.Message.Chat.UserName, chatId)
			}
			if update.Message.Command() == "list" {
				actions.getAllTeams(bot, chatId)
			}
			if update.Message.Command() == "myTeams" {
				actions.getMyTeams(bot, chatId)
			}
			if update.Message.Command() == "subscribe" {
				actions.subscribeOnTeam(bot, updates, chatId)
			}
			if update.Message.Command() == "unsubscribe" {
				actions.unsubscribeFromTeam(bot, updates, chatId)
			}
		}
	}
}

func (actions TelegramActions) subscribeOnTeam(bot *TelegramBot, updates tgbotapi.UpdatesChannel, chatId int64) {
	userId, err := actions.userService.GetUser(int32(chatId))
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

	actions.subscriptionService.SubscribeOnTeam(userId.Id, int32(teamId))

	message(bot, chatId, "Subscribed")
}

func (actions TelegramActions) unsubscribeFromTeam(bot *TelegramBot, updates tgbotapi.UpdatesChannel, chatId int64) {
	userId, err := actions.userService.GetUser(int32(chatId))
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

	actions.subscriptionService.UnubscribeFromTeam(userId.Id, int32(teamId))

	message(bot, chatId, "Unsubscribed")
}

func (actions TelegramActions) getMyTeams(bot *TelegramBot, chatId int64) {
	userId, err := actions.userService.GetUser(int32(chatId))
	if nil != err {
		message(bot, chatId, "Wrong user")
		return
	}

	teams := actions.teamService.GetMyTeams(userId.Id)
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

func (actions TelegramActions) getAllTeams(bot *TelegramBot, chatId int64) {
	teams := actions.teamService.GetTeams()
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

func (actions TelegramActions) saveUser(username string, chatId int64) {
	actions.userService.CreateUser(username, int32(chatId))
}
