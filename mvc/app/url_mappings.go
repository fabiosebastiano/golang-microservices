package app

import (
	"github.com/fabiosebastiano/golang-microservices/mvc/controllers"
)

func mapURL() {
	router.GET("/user/:user_id", controllers.GetUser)

}
