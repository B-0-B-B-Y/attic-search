package main

import (
	"github.com/gin-gonic/gin"

	"attic-search/middleware"
	"attic-search/routes"
)

func main() {
	router := gin.Default()

	router.Use(middleware.SetCORS())
	router.Use(middleware.VerifyUser())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "Up",
		})
	})

	router.GET("/search/:keyword", routes.Search)
	router.POST("/insert", routes.Insert)
	router.POST("/update", routes.Update)
	router.Run()
}
