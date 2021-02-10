package middleware

import "github.com/gin-gonic/gin"

// SetCORS : Allow cross origin sharing on specified hosts
func SetCORS(allowedHosts []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, host := range allowedHosts {
			c.Writer.Header().Set("Access-Control-Allow-Origin", host)
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
