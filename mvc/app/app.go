package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

//StartApplication is the entry point of the app
func StartApplication() {
	mapURL()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}

}
