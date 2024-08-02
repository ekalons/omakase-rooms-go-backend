package handlers

import (
	"net/http"

	"github.com/ekalons/omakase-rooms-go-backend/db"
	"github.com/gin-gonic/gin"
)

func GetRooms(c *gin.Context) {
	rooms, err := db.FetchAllRooms()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch rooms"})
		return
	}

	c.IndentedJSON(http.StatusOK, rooms)
}
