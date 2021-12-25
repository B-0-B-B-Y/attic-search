package routes

import (
	"github.com/gin-gonic/gin"

	"attic-search/database"
)

// SearchGET : Performs a lookup in the database using user-defined keyword
func SearchGET(c *gin.Context) {
	keyword := c.Param("keyword")
	if keyword == "" {
		c.JSON(400, gin.H{
			"Error": "You need to specify a search keyword",
		})
	}

	items, err := database.GetObjects(keyword)
	if err != nil {
		c.JSON(404, gin.H{
			"Error": "No item found",
		})
	}

	c.JSON(200, gin.H{
		"items": items,
	})
}

// FuzzySearchGET : Performs a fuzzy text search against items in the database using user-defined keyword
func FuzzySearchGET(c *gin.Context) {
	keyword := c.Param("keyword")
	if keyword == "" {
		c.JSON(400, gin.H{
			"Error": "You need to specify a search keyword",
		})
	}

	items, err := database.GetFuzzyObjects(keyword)
	if err != nil {
		c.JSON(404, gin.H{
			"Error": "No item found",
		})
	}

	c.JSON(200, gin.H{
		"items": items,
	})
}
