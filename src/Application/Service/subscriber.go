package service

import (
	notification_repository "go-football/src/Domain/Notification/Repository"
	subscription "go-football/src/Domain/Subscription/Model"
	repository "go-football/src/Domain/Subscription/Repository"
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

func (svc SubscriptionService) SubscribeOnTeam(userId, teamId int32) (*subscription.Subscription, error) {
	subscription := subscription.Subscription{UserId: userId, TeamId: teamId}

	result, err := svc.repository.Add(&subscription)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (svc SubscriptionService) UnubscribeFromTeam(userId, teamId int32) error {
	subscription := subscription.Subscription{UserId: userId, TeamId: teamId}

	err := svc.repository.Delete(&subscription)
	if err != nil {
		return err
	}

	return nil
}

func (svc SubscriptionService) GetUserSubscriptions(userId int32) ([]*subscription.Subscription, error) {
	result, err := svc.repository.FindByUser(userId)
	if err != nil {
		return nil, err
	}

	for _, item := range result {
		item.Notification, _ = svc.notificationRepository.FindBySubscription(item.Id)
	}

	return result, nil
}
