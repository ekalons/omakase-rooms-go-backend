package handlers

import (
	"fmt"
	"net/http"

	"github.com/ekalons/omakase-rooms-go-backend/db"
	"github.com/gin-gonic/gin"
)

// getRooms responds with the list of all rooms as JSON.
func GetRooms(c *gin.Context) {
	rooms, err := db.FetchAllRooms()

	if err != nil {
		fmt.Println("Error fetching rooms:", err) // Log the error for debugging
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch rooms"})
		return
	}

	c.IndentedJSON(http.StatusOK, rooms)
}
