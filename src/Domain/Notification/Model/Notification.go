package model

const (
	LiveEventType  = "LIVE_EVENT"
	StartEventType = "START_EVENT"
)

type NotificationTypes [2]string

type Notification struct {
	Id             int32
	SubscriptionId int32
	Type           string
}

func GetNotificationTypes() NotificationTypes {
	return NotificationTypes{LiveEventType, StartEventType}
}
