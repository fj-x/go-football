package service

import (
	user "go-football/src/Domain/User/Model"
	repository "go-football/src/Domain/User/Repository"
)

type UserService struct {
	repository repository.UserRepositoryInterface
}

func NewUserService(repository repository.UserRepositoryInterface) *UserService {
	return &UserService{repository: repository}
}

func (svc UserService) CreateUser(userName string, remoteId int32) (*user.User, error) {
	userExist, err := svc.repository.IsUserExist(remoteId)
	if err != nil {
		return nil, err
	}

	user := user.User{Name: userName, RemoteId: remoteId}
	if true == userExist {
		return &user, nil
	}

	result, err := svc.repository.Add(&user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (svc UserService) GetUser(remoteId int32) (*user.User, error) {
	return svc.repository.GetUser(remoteId)
}
