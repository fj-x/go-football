package service

import (
	subscription "go-football/Domain/Subscription"
	infrastructure "go-football/Infrastructure"
	repository "go-football/Infrastructure/Repository/Subscription"
	"log"
)

func Subscribe(userId, teamId int32) *subscription.Subscription {
	db := infrastructure.MakeMySql()
	repository := repository.New(db)

	subscription := subscription.Subscription{UserId: userId, TeamId: teamId}

	result, err := repository.Add(&subscription)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}
