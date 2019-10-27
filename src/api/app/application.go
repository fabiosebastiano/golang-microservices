package app

import (
	//"github.com/fabiosebastiano/golang-microservices/src/api/log/log_logrus"
	"github.com/fabiosebastiano/golang-microservices/src/api/log/log_zap"
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
	log_zap.Info("Starting application ... loading URLs mapping") //log_zap.Field("step", 1),
	//log_zap.Field("status", "pending"),

	//log_logrus.Info("Starting application ... loading URLs mapping", "step:1", "status:pending")
	mapUrls()
	log_zap.Info("URLs successfully loaded") //log_zap.Field("step", 2),
	//log_zap.Field("status", "success"),

	//log_logrus.Info("URLs successfully loaded", "step:2", "status:success")

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}

}
