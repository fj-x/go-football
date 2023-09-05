package service

import (
	team "go-football/src/Domain/Team"
	infrastructure "go-football/src/Infrastructure"
	repository "go-football/src/Infrastructure/Repository/Team"
	footballdataapi "go-football/src/Infrastructure/Service/footballDataApi"
	"log"
)

// get teams list from db if empty - call request and populate db
func GetTeams() []*team.Team {
	db := infrastructure.MakeMySql()
	repository := repository.New(db)

	result, err := repository.FindAll()
	if err != nil {
		log.Fatalln(err)
	}
	if len(result) == 0 {
		teams := callApi()
		for _, item := range teams {
			repository.Add(item)
		}

		return teams
	}

	return result
}

func GetMyTeams(userId int32) []*team.Team {
	db := infrastructure.MakeMySql()
	repository := repository.New(db)

	result, err := repository.FindAll()
	if err != nil {
		log.Fatalln(err)
	}
	if len(result) == 0 {
		teams := callApi()
		for _, item := range teams {
			repository.Add(item)
		}

		return teams
	}

	return result
}

func callApi() []*team.Team {
	client := footballdataapi.NewClient()
	result, err := client.GetMatchesList()
	if err != nil {
		log.Fatalln(err)
	}

	teams := make([]*team.Team, 0)

	for _, match := range result.Matches {

		item1 := new(team.Team)
		item2 := new(team.Team)

		item1.RemoteId = match.HomeTeam.Id
		item1.Name = match.HomeTeam.Name

		item2.RemoteId = match.AwayTeam.Id
		item2.Name = match.AwayTeam.Name

		teams = append(teams, item1, item2)
	}

	return teams
}
