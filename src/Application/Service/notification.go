package service

import (
	notification "go-football/src/Domain/Notification/Model"
	repository "go-football/src/Domain/Notification/Repository"
)

type NotificationService struct {
	repository repository.NotificationRepositoryInterface
}

func NewNotificationService(repository repository.NotificationRepositoryInterface) *NotificationService {
	return &NotificationService{repository: repository}
}

func (svc NotificationService) GetNotificationTypeList() notification.NotificationTypes {
	return notification.GetNotificationTypes()
}

func (svc *NotificationService) SubscribeOnNotification(subscriptionId int32, notificationType string) (*notification.Notification, error) {
	notification := notification.Notification{SubscriptionId: subscriptionId, Type: notificationType}

	result, err := svc.repository.Add(&notification)
	if err != nil {
		return nil, err
	}

	return result, nil
}
