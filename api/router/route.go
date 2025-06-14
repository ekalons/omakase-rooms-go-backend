package route

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ekalons/omakase-rooms-go-backend/api/handlers"
	"github.com/ekalons/omakase-rooms-go-backend/api/middleware"
	configuration "github.com/ekalons/omakase-rooms-go-backend/config"
)

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":      "healthy",
		"environment": configuration.Cfg.Environment,
		"timestamp":   fmt.Sprintf("%d", time.Now().Unix()),
	})
}

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

	router.GET("/health", healthCheck)

	api := router.Group("/api")
	{
		api.GET("/rooms", handlers.GetRooms)
		api.GET("/room/:id", handlers.GetRoomByID)
		api.GET("/token", handlers.GetToken)
		api.POST("/createRoom", middleware.AuthRequired(), handlers.CreateRoom)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(fmt.Sprintf(":%s", port))
}
