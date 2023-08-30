package service

import (
	user "go-football/Domain/User"
	infrastructure "go-football/Infrastructure"
	repository "go-football/Infrastructure/Repository/User"
	"log"
)

func CreateUser(userName string) *user.User {
	db := infrastructure.MakeMySql()
	repository := repository.New(db)

	user := user.User{Name: userName}

	result, err := repository.Add(&user)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}
