package services

import (
	"net/http"
	"testing"

	"github.com/fabiosebastiano/golang-microservices/mvc/model"
	"github.com/fabiosebastiano/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(userId int64)(*model.User, *utils.ApplicationError) 
)

func init(){
	model.UserDao = &userDaoMock
}

type usersDaoMock struct {
}

func (m *usersDaoMock) GetUser(userId int64) (*model.User, *utils.ApplicationError) {
	return getUserFunction(userId)
	//return &model.User{}, nil
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64)(*model.User, *utils.ApplicationError){
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Meesage: "user 0 does not exists"
		}
	}

	user, error := UserService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, error)
	assert.EqualValues(t, http.StatusNotFound, error.StatusCode)
	assert.EqualValues(t, "User with id 0 not found", error.Message)

}
func TestGetUserNoError(t *testing.T) {

	getUserFunction = func(userId int64)(*model.User, *utils.ApplicationError){
		return &model.User{
			userId: 1,
			FirstName: "fabio",
		}, nil
	}

	user, error := UserService.GetUser(1)

	assert.Nil(t, error)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "fabio", user.FirstName)

}
