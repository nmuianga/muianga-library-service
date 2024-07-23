package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nmuianga/muianga-library-service/resources"
)

func HealthRouter(r *gin.Engine) {
	r.GET("/health", resources.Health)
}
