package main

import (
	"bufio"
	// "go-football/src/Infrastructure/Service/telegram"
	serviceOp "go-football/src/Application/Service"
	infrastructure "go-football/src/Infrastructure"
	notification_repository "go-football/src/Infrastructure/Repository/Notification"
	subscription_repository "go-football/src/Infrastructure/Repository/Subscription"
	team_repository "go-football/src/Infrastructure/Repository/Team"
	user_repository "go-football/src/Infrastructure/Repository/User"

	"os"

	"github.com/joho/godotenv"
)

// Load env data.
func init() {
	godotenv.Load()
}

var in = bufio.NewReader(os.Stdin)

func main() {

	// init db
	db := infrastructure.MakeMySql()

	// init repositories
	notificationRepository := notification_repository.New(db)
	subscriptionRepository := subscription_repository.New(db)
	userRepository := user_repository.New(db)
	teamRepository := team_repository.New(db)

	// init services
	serviceOp.NewUserService(userRepository)
	serviceOp.NewTeamService(teamRepository)
	serviceOp.NewNotificationService(notificationRepository)
	serviceOp.NewSubscriptionService(subscriptionRepository, notificationRepository)

	// Initialise scheduler
	serviceOp.StartMatchScheduler()

	// Run actions
	// telegram.Actions()

	//serviceOp.GetTeams()
	// ntf := serviceOp.CreateUser("MyUser")
	// ntf := serviceOp.SubscribeOnTeam(1, 3)
	// ntf := serviceOp.SubscribeOnNotification(1, "START_EVENT")

	// ntf := serviceOp.GetNotificationTypeList()
	// fmt.Println(ntf)
	// serviceOp.GetTeams("PL")
	// loop:
	// 	for {
	// 		fmt.Println("Select command")

	// 		choice, _ := in.ReadString('\n')

	//		switch strings.TrimSpace(choice) {
	//		case "1":
	//			serviceOp.GetTeams("PL")
	//		case "0":
	//			break loop
	//		default:
	//			fmt.Println("unknown")
	//		}
	//	}
}
