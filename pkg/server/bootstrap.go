package server

import (
	"log"
	config "simpleService/internal/config"
	"simpleService/internal/config/errorcode"
	"simpleService/internal/pg"
	"simpleService/pkg/route"

	"github.com/labstack/echo/v4"
)

// Bootstrap ...
func Bootstrap(e *echo.Echo) string {
	// base
	config.Init()

	// pg connect
	err := pg.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// errorcode init
	errorcode.Init()

	// route init
	route.Init(e)

	// services
	cfg := config.GetENV()
	return cfg.Port
}
