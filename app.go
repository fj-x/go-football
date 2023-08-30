package main

import (
	"bufio"
	"fmt"
	serviceOp "go-football/src/Application/Service"
	"os"

	"github.com/joho/godotenv"
)

// Load env data.
func init() {
	godotenv.Load()
}

var in = bufio.NewReader(os.Stdin)

func main() {

	// ntf := serviceOp.GetTeams("PL")
	// ntf := serviceOp.CreateUser("MyUser")
	// ntf := serviceOp.SubscribeOnTeam(1, 3)
	// ntf := serviceOp.SubscribeOnNotification(1, "START_EVENT")

	ntf := serviceOp.GetNotificationTypeList()
	fmt.Println(ntf)
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
