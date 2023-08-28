package service

import (
	team "go-football/Domain/Team"
	infrastructure "go-football/Infrastructure"
	repository "go-football/Infrastructure/Repository"
	footballdataapi "go-football/Infrastructure/Service/footballDataApi"
	"log"
)

// get teams list from db if empty call request and populate db
func GetTeams(league string) []team.Team {
	db := infrastructure.MakeMySql()
	repository := repository.New(db)

	result, err := repository.FindAll()
	if err != nil {
		log.Fatalln(err)
	}
	if result != nil {
		teams := callApi(league)
		for _, item := range teams {
			repository.Add(&item)
		}

		return teams
	}

	return result
}

func callApi(league string) []team.Team {
	client := footballdataapi.NewClient()
	result, err := client.GetMatchesList(league)
	if err != nil {
		log.Fatalln(err)
	}

	var teams []team.Team
	for _, match := range result.Matches {

		teams = append(teams,
			team.Team{
				Id:   match.HomeTeam.Id,
				Name: match.HomeTeam.Name,
			}, team.Team{
				Id:   match.AwayTeam.Id,
				Name: match.AwayTeam.Name,
			})
	}

	return teams
}
