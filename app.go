package main

import (
	"bufio"
	"os"

	"github.com/joho/godotenv"
)

// Load env data.
func init() {
	godotenv.Load()
}

var in = bufio.NewReader(os.Stdin)

func main() {

	// serviceOp.Subscribe(1, 42)
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
