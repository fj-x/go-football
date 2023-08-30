package service

import (
	subscription "go-football/src/Domain/Subscription"
	infrastructure "go-football/src/Infrastructure"
	repository "go-football/src/Infrastructure/Repository/Subscription"
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
