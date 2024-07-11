package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome")
	/*r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Live",
		})
	})
	r.Run()*/
	fmt.Printf("Sum: %d", sum(3, 5))
}

func sum(a int, b int) int {
	return a + b
}
