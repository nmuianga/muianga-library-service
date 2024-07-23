package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nmuianga/muianga-library-service/utils/db"
)

func main() {
	db.OpenConnection()
	fmt.Println("Welcome")
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Live",
		})
	})
	r.Run()

}
