package main

import (
	"os"
	"simpleService/pkg/server"

	config "simpleService/internal/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title       simpleService
// @version     1.0
// @description API for simpleService
// @description
// @description ******************************
// @description - Add description
// @description ******************************
// @description
// @basePath       /simple-service

// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization
func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))

	e.Use(middleware.Gzip())
	if os.Getenv("ENV") == "release" {
		e.Use(middleware.Recover())
	}

	// Bootstrap things
	server.Bootstrap(e)

	// Start server
	e.Logger.Fatal(e.Start(config.GetENV().Port))
}
