package service

import (
	team "go-football/src/Domain/Team/Model"
	repository "go-football/src/Domain/Team/Repository"
	footballdataapi "go-football/src/Infrastructure/Service/footballDataApi"
	"log"
)

type TeamService struct {
	repository repository.TeamRepositoryInterface
}

func NewTeamService(repository repository.TeamRepositoryInterface) *TeamService {
	return &TeamService{repository: repository}
}

// get teams list from db if empty - call request and populate db
func (svc TeamService) GetTeams() []*team.Team {
	result, err := svc.repository.FindAll()
	if err != nil {
		log.Fatalln(err)
	}
	if len(result) == 0 {
		teams := callApi()
		for _, item := range teams {
			svc.repository.Add(item)
		}

		return teams
	}

	return result
}

func (svc TeamService) GetMyTeams(userId int32) []*team.Team {
	result, err := svc.repository.FindUsersTeams(userId)
	if err != nil {
		log.Fatalln(err)
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
