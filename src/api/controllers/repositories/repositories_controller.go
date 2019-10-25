package repositories

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fabiosebastiano/golang-microservices/src/api/domain/repositories"
	"github.com/fabiosebastiano/golang-microservices/src/api/services"
	"github.com/fabiosebastiano/golang-microservices/src/api/utils/errors"
)

// CreateRepo .
func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiError := errors.NewBadRequestError("invalid json body")
		c.JSON(apiError.Status(), apiError)
		return
	}

	result, error := services.RepositoryService.CreateRepo(request)
	if error != nil {
		c.JSON(error.Status(), error)
		return
	}
	c.JSON(http.StatusCreated, result)

}
