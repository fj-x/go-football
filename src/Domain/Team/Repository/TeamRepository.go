package repository

import (
	team "go-football/src/Domain/Team/Model"
)

type TeamRepositoryInterface interface {
	FindAll() ([]*team.Team, error)
	FindUsersTeams(userId int32) ([]*team.Team, error)
	Add(item *team.Team) (int64, error)
}
