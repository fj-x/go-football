package service

import (
	team "go-football/Domain/Team"
	infrastructure "go-football/Infrastructure"
	repository "go-football/Infrastructure/Repository"
	"log"
)

// get teams list from db if empty - call request and populate db
func Subscribe(userId, teamId int32) []*team.Team {
	db := infrastructure.MakeMySql()
	repository := repository.New(db)

	result, err := repository.FindAll()
	if err != nil {
		log.Fatalln(err)
	}

	return result
}
