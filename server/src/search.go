package main

import (
	database "./database"
	"github.com/gin-gonic/gin"
)

// searchGET : Performs a lookup in the database using user-defined keyword
func searchGET(c *gin.Context) {
	keyword := c.Param("keyword")
	if keyword == "" {
		c.JSON(400, gin.H{
			"Error": "You need to specify a search keyword",
		})
	}

	items, err := database.GetObject(keyword)
	if err != nil {
		c.JSON(404, gin.H{
			"Error": "No item found",
		})
	}

	c.JSON(200, gin.H{
		"items": items,
	})
}
