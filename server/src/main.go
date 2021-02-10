package main

import (
	"github.com/gin-gonic/gin"

	"attic-search/middleware"
	"attic-search/routes"
)

func main() {
	allowedHosts := []string{
		"http://localhost",
		"http://localhost:19006",
	}
	router := gin.Default()

	router.Use(middleware.SetCORS(allowedHosts))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "Up",
		})
	})

	router.GET("/search/word/:keyword", routes.SearchGET)
	router.GET("/search/fuzzy/:keyword", routes.FuzzySearchGET)
	router.Run(":3000")
}
