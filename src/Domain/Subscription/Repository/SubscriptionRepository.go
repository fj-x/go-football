package repository

import (
	subscription "go-football/src/Domain/Subscription/Model"
)

type SubscriptionRepositoryInterface interface {
	FindAll() ([]*subscription.Subscription, error)
	FindUnqueSubscribedTeams() ([]int32, error)
	FindMatchSubscribers(homeTeam, awayTeam int32) ([]int32, error)
	FindByUser(userId int32) ([]*subscription.Subscription, error)
	Add(item *subscription.Subscription) (*subscription.Subscription, error)
	Delete(item *subscription.Subscription) error
}
