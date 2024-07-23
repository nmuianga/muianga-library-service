package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nmuianga/muianga-library-service/routers"
)

func main() {
	r := gin.Default()
	routers.HealthRouter(r)
	r.Run()

}
