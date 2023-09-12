package service

import (
	"fmt"
	repository "go-football/src/Domain/Subscription/Repository"

	footballdataapi "go-football/src/Infrastructure/Service/footballDataApi"
	"time"

	"log"
	// "github.com/robfig/cron/v3"
)

type MatchSchedulerService struct {
	repository repository.SubscriptionRepositoryInterface
}

func NewSchedulerService(
	repository repository.SubscriptionRepositoryInterface,
) *MatchSchedulerService {
	return &MatchSchedulerService{repository: repository}
}

var monitoringList []int32

func (svc MatchSchedulerService) StartMatchScheduler() {
	// c := cron.New()
	// c.AddFunc("00 00 * * *", shceduleMatches)
	// c.Start()

	svc.shceduleMatches()
}

func (svc MatchSchedulerService) shceduleMatches() {
	monitoringList, err := svc.repository.FindUnqueSubscribedTeams()
	if err != nil {
		log.Fatalln(err)
	}

	// call api to get all matches
	client := footballdataapi.NewClient()
	nextMatches, err := client.GetMatchesList()
	if err != nil {
		log.Fatalln(err)
	}

	for _, match := range nextMatches.Matches {
		if inSlice(match.HomeTeam.Id, monitoringList) || inSlice(match.AwayTeam.Id, monitoringList) {
			// run goroutine
			// go monitorLiveMatch(match)
			svc.monitorLiveMatch(match)
		}
	}
}

func (svc MatchSchedulerService) monitorLiveMatch(match footballdataapi.Match) {
	date, error := time.Parse("2006-01-02T15:04:05Z", match.StartDate)
	if error != nil {
		fmt.Println(error)
		return
	}

	// Create a timer to start monitoring at matchStartTime
	timer := time.NewTimer(date.Sub(time.Now()))

	// Wait for the timer to expire
	<-timer.C

	var goalsMap = make(map[int32][]footballdataapi.Goal)

	client := footballdataapi.NewClientMock()
	// Continue monitoring the match
	for {
		// Fetch match information from the API
		// call api

		matchInfo, err := client.FetchMatchInfo(match.Id)
		if err != nil || matchInfo == nil {
			// Handle error
			fmt.Println(err, "fuuuuuu")
			return
		}

		// Check if the match has ended
		if "FINISHED" == matchInfo.Status {
			// Exit the goroutine
			return
		}

		newGoals := findSliceDifferences(matchInfo.Goals, goalsMap[matchInfo.Id])

		for _, newGoal := range newGoals {
			goalsMap[matchInfo.Id] = append(goalsMap[matchInfo.Id], newGoal)
			svc.notify(newGoal, match)
		}

		// Sleep for a minute before checking again
		time.Sleep(60 * time.Second)
	}
}

func (svc MatchSchedulerService) notify(goal footballdataapi.Goal, match footballdataapi.Match) {
	subscribers, err := svc.repository.FindMatchSubscribers(match.HomeTeam.Id, match.AwayTeam.Id)
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
