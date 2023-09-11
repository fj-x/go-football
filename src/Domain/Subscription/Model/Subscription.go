package model

import notification "go-football/src/Domain/Notification/Model"

type Subscription struct {
	Id           int32
	UserId       int32
	TeamId       int32
	Notification []*notification.Notification
}
