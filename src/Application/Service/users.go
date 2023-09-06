package service

import (
	user "go-football/src/Domain/User"
	infrastructure "go-football/src/Infrastructure"
	repository "go-football/src/Infrastructure/Repository/User"
	"log"
)

func CreateUser(userName string, remoteId int32) *user.User {
	db := infrastructure.MakeMySql()
	repository := repository.New(db)

	userExist, _ := repository.IsUserExist(remoteId)

	user := user.User{Name: userName, RemoteId: remoteId}
	if true == userExist {
		return &user
	}

	result, err := repository.Add(&user)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}

func GetUser(remoteId int32) (*user.User, error) {
	db := infrastructure.MakeMySql()
	repository := repository.New(db)

	return repository.GetUser(remoteId)
}
