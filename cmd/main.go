package main

import (
	route "github.com/ekalons/omakase-rooms-go-backend/api/router"
	"github.com/ekalons/omakase-rooms-go-backend/db"

	configuration "github.com/ekalons/omakase-rooms-go-backend/config"
)

func main() {
	configuration.Load()

	db.Connect()
	defer db.Disconnect()

	route.Setup()
}
