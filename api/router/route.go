package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ekalons/omakase-rooms-go-backend/api/handlers"
	"github.com/ekalons/omakase-rooms-go-backend/api/middleware"
	configuration "github.com/ekalons/omakase-rooms-go-backend/config"
)

func Setup() {

	router := gin.Default()
	if configuration.Cfg.Environment == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{configuration.Cfg.FrontEndUrl},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/rooms", handlers.GetRooms)
	router.GET("/room/:id", handlers.GetRoomByID)
	router.GET("/token", handlers.GetToken)
	router.POST("/createRoom", middleware.AuthRequired(), handlers.CreateRoom)

	router.Run(":8080")
}
