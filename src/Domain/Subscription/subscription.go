package subscription

import notification "go-football/src/Domain/Notification"

type Subscription struct {
	Id           int32
	UserId       int32
	TeamId       int32
	Notification []*notification.Notification
}
