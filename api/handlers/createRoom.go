package handlers

import (
	"net/http"

	"github.com/ekalons/omakase-rooms-go-backend/db"
	"github.com/ekalons/omakase-rooms-go-backend/models"
	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context) {
	var newRoom models.Room

	if err := c.BindJSON(&newRoom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result, err := db.InsertRoom(newRoom)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"insertedID": result.InsertedID,
		"room":       newRoom,
	})
}
