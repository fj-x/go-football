package service

import (
	user "go-football/src/Domain/User/Model"
	repository "go-football/src/Domain/User/Repository"
	"log"
)

type UserService struct {
	repository repository.UserRepositoryInterface
}

func NewUserService(repository repository.UserRepositoryInterface) *UserService {
	return &UserService{repository: repository}
}

func (svc UserService) CreateUser(userName string, remoteId int32) *user.User {
	userExist, _ := svc.repository.IsUserExist(remoteId)

	user := user.User{Name: userName, RemoteId: remoteId}
	if true == userExist {
		return &user
	}

	result, err := svc.repository.Add(&user)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}

func (svc UserService) GetUser(remoteId int32) (*user.User, error) {
	return svc.repository.GetUser(remoteId)
}
