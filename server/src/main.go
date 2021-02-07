package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "Up",
		})
	})

	router.GET("/search/:keyword", searchGET)
	router.Run(":3000")
}
