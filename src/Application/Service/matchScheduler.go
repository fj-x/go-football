package service

import (
	"fmt"
	infrastructure "go-football/src/Infrastructure"
	repository "go-football/src/Infrastructure/Repository/Subscription"
	footballdataapi "go-football/src/Infrastructure/Service/footballDataApi"
	"time"

	"log"

	"github.com/robfig/cron/v3"
)

var monitoringList []int32

func StartMatchScheduler() {
	c := cron.New()
	c.AddFunc("00 00 * * *", shceduleMatches)
	c.Start()
}

func shceduleMatches() {
	db := infrastructure.MakeMySql()
	repository := repository.New(db)

	monitoringList, err := repository.FindUnqueSubscribedTeams()
	if err != nil {
		log.Fatalln(err)
	}

	// call api
	client := footballdataapi.NewClient()
	nextMatches, err := client.GetMatchesList()
	if err != nil {
		log.Fatalln(err)
	}

	for _, match := range nextMatches.Matches {
		if inSlice(match.HomeTeam.Id, monitoringList) || inSlice(match.AwayTeam.Id, monitoringList) {
			// run goroutine
			go monitorLiveMatch(match)
		}
	}
}

func monitorLiveMatch(match footballdataapi.Match) {
	date, error := time.Parse("2023-09-01T19:00:00Z", match.StartDate)
	if error != nil {
		fmt.Println(error)
		return
	}

	// Create a timer to start monitoring at matchStartTime
	timer := time.NewTimer(date.Sub(time.Now()))

	// Wait for the timer to expire
	<-timer.C

	var goals map[int32][]footballdataapi.Goal

	// Continue monitoring the match
	for {
		// Fetch match information from the API
		// call api
		client := footballdataapi.NewClient()
		matchInfo, err := client.FetchMatchInfo(match.Id)
		if err != nil {
			// Handle error
		}

		// Check if the match has ended
		if "FINISHED" == matchInfo.Status {
			// Exit the goroutine
			return
		}

		newGoals := findSliceDifferences(matchInfo.Goals, goals[matchInfo.Id])
		for _, newGoal := range newGoals {
			goals[matchInfo.Id] = append(goals[matchInfo.Id], newGoal)
			notify(newGoal, match)
		}

		// Sleep for a minute before checking again
		time.Sleep(60 * time.Second)
	}
}

func notify(goal footballdataapi.Goal, match footballdataapi.Match) {
	db := infrastructure.MakeMySql()
	repository := repository.New(db)

	subscribers, err := repository.FindMatchSubscribers(match)
	if err != nil {
		log.Fatalln(err)
	}

	// Check for score changes and notify subscribed users
	for _, user := range subscribers {
		fmt.Println(user, "Score has changed!")
	}
}

func inSlice(a int32, list []int32) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func findSliceDifferences(slice1, slice2 []footballdataapi.Goal) []footballdataapi.Goal {
	// Create a map to store the unique elements from both slices
	uniqueElements := make(map[footballdataapi.Goal]bool)

	// Iterate over the first slice and add its elements to the map
	for _, s := range slice1 {
		uniqueElements[s] = true
	}

	// Iterate over the second slice and remove elements found in the first slice
	for _, s := range slice2 {
		delete(uniqueElements, s)
	}

	// Collect the unique elements that were found in the first slice but not in the second
	differences := []footballdataapi.Goal{}
	for s := range uniqueElements {
		differences = append(differences, s)
	}

	return differences
}