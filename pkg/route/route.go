package route

import (
	"github.com/labstack/echo/v4"
)

// Init ...
func Init(e *echo.Echo) {
	g := e.Group("/simple-service")

	// Components
	user(g)
}
