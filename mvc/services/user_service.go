package services

import (
	"github.com/fabiosebastiano/golang-microservices/mvc/model"
	"github.com/fabiosebastiano/golang-microservices/mvc/utils"
)

type userService struct {
}

var (
	UserService userService
)

//GetUser .
func (u *userService) GetUser(userID int64) (*model.User, *utils.ApplicationError) {
	user, err := model.UserDao.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
