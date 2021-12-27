package routes

import (
	"github.com/gin-gonic/gin"

	"attic-search/database"
)

// Perform a fuzzy word search across the DB and return any items that contain the keyword in any of their fields
func Search(c *gin.Context) {
	keyword := c.Param("keyword")
	if keyword == "" {
		c.JSON(400, gin.H{
			"Error": "You need to specify a search keyword",
		})
		return
	}

	result, err := database.SearchForItem(keyword)
	if err != nil {
		status := 500
		if err.Error() == "record does not exist" {
			status = 404
		}

		c.JSON(status, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"items": result,
	})
}
