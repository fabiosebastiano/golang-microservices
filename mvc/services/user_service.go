package services

import (
	"github.com/fabiosebastiano/golang-microservices/mvc/model"
	"github.com/fabiosebastiano/golang-microservices/mvc/utils"
)

//GetUser .
func GetUser(userID int64) (*model.User, *utils.ApplicationError) {
	return model.GetUser(userID)
}
