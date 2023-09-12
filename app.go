package main

import (
	service "go-football/src/Application/Service"
	infrastructure "go-football/src/Infrastructure"
	notification_repository "go-football/src/Infrastructure/Repository/Notification"
	subscription_repository "go-football/src/Infrastructure/Repository/Subscription"
	team_repository "go-football/src/Infrastructure/Repository/Team"
	user_repository "go-football/src/Infrastructure/Repository/User"
	telegram "go-football/src/Infrastructure/Service/telegram"

	"github.com/joho/godotenv"
)

// Load env data.
func init() {
	godotenv.Load()
}

func main() {

	// init db
	db := infrastructure.MakeMySql()

	// init repositories
	notificationRepository := notification_repository.New(db)
	subscriptionRepository := subscription_repository.New(db)
	userRepository := user_repository.New(db)
	teamRepository := team_repository.New(db)

	// init services
	userService := service.NewUserService(userRepository)
	teamService := service.NewTeamService(teamRepository)
	notificationService := service.NewNotificationService(notificationRepository)
	subscriptionService := service.NewSubscriptionService(subscriptionRepository, notificationRepository)

	// Initialise scheduler
	sheduler := service.NewSchedulerService(subscriptionRepository)
	sheduler.StartMatchScheduler()

	// Run actions
	telegram.New(userService, teamService, subscriptionService, notificationService).Actions()
}
