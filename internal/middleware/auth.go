package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/irawankilmer/lms_backend/internal/config"
	"github.com/irawankilmer/lms_backend/internal/models"
)

func ClientAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := c.GetHeader("X-Client-ID")
		clientSecret := c.GetHeader("X-Client-Secret")

		var client models.Client
		result := config.DB.Where("client_id = ? AND client_secret = ?", clientID, clientSecret).First(&client)

		if result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
