package repository

import (
	user "go-football/src/Domain/User/Model"
)

type UserRepositoryInterface interface {
	FindAll() ([]*user.User, error)
	GetUser(remoteId int32) (*user.User, error)
	Add(item *user.User) (*user.User, error)
	IsUserExist(userId int32) (bool, error)
}
