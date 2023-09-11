package service

import (
	notification_repository "go-football/src/Domain/Notification/Repository"
	subscription "go-football/src/Domain/Subscription/Model"
	repository "go-football/src/Domain/Subscription/Repository"
	"log"
)

type SubscriptionService struct {
	repository             repository.SubscriptionRepositoryInterface
	notificationRepository notification_repository.NotificationRepositoryInterface
}

func NewSubscriptionService(
	repository repository.SubscriptionRepositoryInterface,
	notificationRepository notification_repository.NotificationRepositoryInterface,
) *SubscriptionService {
	return &SubscriptionService{repository: repository, notificationRepository: notificationRepository}
}

func (svc SubscriptionService) SubscribeOnTeam(userId, teamId int32) *subscription.Subscription {
	subscription := subscription.Subscription{UserId: userId, TeamId: teamId}

	result, err := svc.repository.Add(&subscription)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}

func (svc SubscriptionService) UnubscribeFromTeam(userId, teamId int32) {
	subscription := subscription.Subscription{UserId: userId, TeamId: teamId}

	err := svc.repository.Delete(&subscription)
	if err != nil {
		log.Fatalln(err)
	}
}

func (svc SubscriptionService) GetUserSubscriptions(userId int32) []*subscription.Subscription {
	result, err := svc.repository.FindByUser(userId)
	if err != nil {
		log.Fatalln(err)
	}

	for _, item := range result {
		item.Notification, _ = svc.notificationRepository.FindBySubscription(item.Id)
	}

	return result
}
