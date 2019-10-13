package services

import (
	"net/http"

	"github.com/fabiosebastiano/golang-microservices/mvc/model"
	"github.com/fabiosebastiano/golang-microservices/mvc/utils"
)

type itemService struct{}

var (
	ItemService itemService
)

//GetItem .
func (i *itemService) GetItem(itemId string) (*model.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "to be implemented",
		StatusCode: http.StatusInternalServerError,
	}
}
