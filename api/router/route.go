package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ekalons/omakase-rooms-go-backend/api/handlers"
	configuration "github.com/ekalons/omakase-rooms-go-backend/config"
)

func Setup() {

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
