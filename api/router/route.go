package route

import (
	"strings"

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
		AllowOrigins:     strings.Split(configuration.Cfg.FrontEndUrl, ","),
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := router.Group("/api")
	{
		api.GET("/rooms", handlers.GetRooms)
		api.GET("/room/:id", handlers.GetRoomByID)
		api.GET("/token", handlers.GetToken)
		api.POST("/createRoom", middleware.AuthRequired(), handlers.CreateRoom)
	}

	router.Run(":8080")
}
