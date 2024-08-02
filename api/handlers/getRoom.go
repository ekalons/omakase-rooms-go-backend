package handlers

import (
	"net/http"

	"github.com/ekalons/omakase-rooms-go-backend/db"
	"github.com/gin-gonic/gin"
)

func GetRoomByID(c *gin.Context) {
	id := c.Param("id")

	room, err := db.FetchRoomById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch room by ID"})
		return
	}

	if room == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Room not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, room)
}
