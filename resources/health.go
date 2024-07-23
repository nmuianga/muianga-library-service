package resources

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Live")
}
