package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/quantinium3/lucy/cmd/handler"
)

func SetupUserRoutes(api *echo.Group, userService *handler.Service) {
	api.POST("/user/create", userService.CreateUser)
	api.DELETE("/user/:id", userService.DeleteUser)
	api.GET("/user/:id", userService.GetUser)
}
