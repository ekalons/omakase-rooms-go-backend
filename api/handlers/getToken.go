package handlers

import (
	"net/http"

	"github.com/ekalons/omakase-rooms-go-backend/api/middleware"
	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) {
	clientSecret := c.GetHeader("X-Client-Secret")

	token, err := middleware.GenerateToken(clientSecret)
	if err != nil {
		if err.Error() == "invalid client secret" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid client secret"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
