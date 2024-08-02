package main

import (
	"github.com/ekalons/omakase-rooms-go-backend/db"
	"github.com/ekalons/omakase-rooms-go-backend/handlers"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	configuration "github.com/ekalons/omakase-rooms-go-backend/config"
)

func main() {
	configuration.Load()

	db.Connect()
	defer db.Disconnect()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{configuration.Cfg.FrontEndUrl},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/rooms", handlers.GetRooms)
	router.GET("/rooms/:id", handlers.GetRoomByID)
	router.POST("/rooms", handlers.PostRooms)

	router.Run("localhost:8080")

}
