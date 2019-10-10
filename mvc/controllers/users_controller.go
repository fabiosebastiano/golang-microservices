package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fabiosebastiano/golang-microservices/mvc/services"
	"github.com/fabiosebastiano/golang-microservices/mvc/utils"
)

//GetUsers controller
func GetUsers(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("tutto è bene quel che finisce bene"))
}

//GetUser controller
func GetUser(response http.ResponseWriter, request *http.Request) {
	//VALIDAZIONE PARAMETRI -> UNICA LOGICA DEL CONTROLLER
	userID, err := strconv.ParseInt(request.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiError := &utils.ApplicationError{
			Message:    "User id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiError)
		response.WriteHeader(apiError.StatusCode)
		response.Write(jsonValue)
		return
	}

	//SE I PARAMETRI SONO OK, ALLORA CHIAMA SERVICE
	user, apiError := services.GetUser(userID)
	if apiError != nil {
		jsonValue, _ := json.Marshal(apiError)
		response.WriteHeader(apiError.StatusCode)
		response.Write(jsonValue)
		return
	}
	jsonvalue, _ := json.Marshal(user)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)
	response.Write(jsonvalue)
}
