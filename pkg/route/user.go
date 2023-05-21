package route

import (
	"simpleService/pkg/handler"

	routevalidation "simpleService/pkg/route/validation"

	"github.com/labstack/echo/v4"
)

func user(e *echo.Group) {
	g := e.Group("/users")
	h := handler.User()
	v := routevalidation.User()

	g.GET("/login", h.Login, v.Login)
	g.POST("/logout", h.Logout, v.Logout)
	g.POST("", h.Register, v.Register)
}
