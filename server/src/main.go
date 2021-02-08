package main

import (
	"github.com/gin-gonic/gin"

	"attic-search/routes"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "Up",
		})
	})

	router.GET("/search/word/:keyword", routes.SearchGET)
	router.GET("/search/fuzzy/:keyword", routes.FuzzySearchGET)
	router.Run(":3000")
}
