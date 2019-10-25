package app

import (
	"github.com/fabiosebastiano/golang-microservices/src/api/controllers/health"
	"github.com/fabiosebastiano/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/health", health.Liveness)
	router.POST("/repositories", repositories.CreateRepo)
}
