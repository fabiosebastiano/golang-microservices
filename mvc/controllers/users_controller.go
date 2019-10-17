package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/fabiosebastiano/golang-microservices/mvc/services"
	"github.com/fabiosebastiano/golang-microservices/mvc/utils"
)

//GetUser controller
func GetUser(c *gin.Context) { //VERSIONE CON GINGONIC

	//VALIDAZIONE PARAMETRI -> UNICA LOGICA DEL CONTROLLER
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiError := &utils.ApplicationError{
			Message:    "User id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.Respond(c, apiError.StatusCode, apiError)
		//RITORNO ERRORE IN FORMATO JSON ATTRAVERSO C
		//c.JSON(apiError.StatusCode, apiError)
		return
	}

	//SE I PARAMETRI SONO OK, ALLORA CHIAMA SERVICE
	user, apiError := services.UserService.GetUser(userID)
	if apiError != nil {
		utils.Respond(c, apiError.StatusCode, apiError)
		return
	}
	utils.Respond(c, http.StatusOK, user)

}

func respond(accept string) {

}
