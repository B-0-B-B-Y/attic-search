package middleware

import (
	"log"
	"path/filepath"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

// SetCORS : Allow cross origin sharing
func SetCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
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

// VerifyUser : Verifies API calls are being made from logged in legitimate Firebase Auth users
func VerifyUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceAccount, _ := filepath.Abs("serviceAccount.json")
		opt := option.WithCredentialsFile(serviceAccount)
		firebaseApp, err := firebase.NewApp(c, nil, opt)
		if err != nil {
			log.Printf("Error initialising the firebase admin SDK. Reason => %v", err)
			c.AbortWithStatus(500)
			return
		}

		firebaseAuth, err := firebaseApp.Auth(c)
		if err != nil {
			log.Printf("Error initialising the firebase auth client. Reason => %v", err)
			c.AbortWithStatus(500)
			return
		}

		_, err = firebaseAuth.VerifyIDToken(c, c.Request.Header.Get("Authorization"))
		if err != nil {
			log.Printf("Error verifying user ID token. Reason => %v", err)
			c.AbortWithStatus(401)
			return
		}

		c.Next()
	}
}
