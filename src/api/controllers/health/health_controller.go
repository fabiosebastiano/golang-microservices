package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	liveness = "I'm Alive!"
)

func Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, liveness)
}
