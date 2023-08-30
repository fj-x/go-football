package service

import (
	subscription "go-football/src/Domain/Subscription"
	infrastructure "go-football/src/Infrastructure"
	notification_repository "go-football/src/Infrastructure/Repository/Notification"
	repository "go-football/src/Infrastructure/Repository/Subscription"
	"log"
)

func SubscribeOnTeam(userId, teamId int32) *subscription.Subscription {
	db := infrastructure.MakeMySql()
	repository := repository.New(db)

	subscription := subscription.Subscription{UserId: userId, TeamId: teamId}

	result, err := repository.Add(&subscription)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}

func GetUserSubscriptions(userId int32) []*subscription.Subscription {
	db := infrastructure.MakeMySql()
	repository := repository.New(db)
	notification_repository := notification_repository.New(db)

	result, err := repository.FindByUser(userId)
	if err != nil {
		log.Fatalln(err)
	}

	for _, item := range result {
		item.Notification, _ = notification_repository.FindBySubscription(item.Id)
	}

	return result
}
