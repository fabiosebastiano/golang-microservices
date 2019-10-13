package model

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fabiosebastiano/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		1: {1, "fabio", "sebastiano", "fs@ciao.it"},
	}
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

//GetUser .
func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {

	log.Println("Getting data from DB")

	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User with id %v not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
