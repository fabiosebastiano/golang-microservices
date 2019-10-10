package model

import (
	"fmt"
	"net/http"

	"github.com/fabiosebastiano/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		1: {1, "fabio", "sebastiano", "fs@ciao.it"},
	}
)

//GetUser .
func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
