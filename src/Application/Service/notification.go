package service

import (
	notification "go-football/src/Domain/Notification"
	infrastructure "go-football/src/Infrastructure"
	repository "go-football/src/Infrastructure/Repository/Notification"
	"log"
)

func GetNotificationTypeList() notification.NotificationTypes {
	return notification.GetNotificationTypes()
}

func SubscribeOnNotification(subscriptionId int32, notificationType string) *notification.Notification {
	db := infrastructure.MakeMySql()
	repository := repository.New(db)

	notification := notification.Notification{SubscriptionId: subscriptionId, Type: notificationType}

	result, err := repository.Add(&notification)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}
