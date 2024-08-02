package handlers

import (
	"net/http"

	"github.com/ekalons/omakase-rooms-go-backend/models"
	"github.com/gin-gonic/gin"
)

func PostRooms(c *gin.Context) {
	var newRoom models.Room

	// Call BindJSON to bind the received JSON to
	// newRoom.
	if err := c.BindJSON(&newRoom); err != nil {
		return
	}

	// Add the new room to the slice.
	// rooms = append(rooms, newRoom)
	c.IndentedJSON(http.StatusCreated, newRoom)
}
