package model

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserUserNotFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "unexpected user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "User with id 0 not found", err.Message)

}

func TestGetUser(t *testing.T) {
	user, err := GetUser(1)

	assert.Nil(t, err, "we weren't expecting an error when user id is 1")
	assert.NotNil(t, user, "we were expecting a user")
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "fabio", user.FirstName)
	assert.EqualValues(t, "sebastiano", user.LastName)
	assert.EqualValues(t, "fs@ciao.it", user.Email)
}
