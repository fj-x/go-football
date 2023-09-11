package repository

import (
	model "go-football/src/Domain/Notification/Model"
)

type NotificationRepositoryInterface interface {
	Add(item *model.Notification) (*model.Notification, error)
	FindBySubscription(subscriptionId int32) ([]*model.Notification, error)
}
