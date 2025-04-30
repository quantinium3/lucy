package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/quantinium3/lucy/cmd/handler"
)

func SetupPeripheralRoutes(api *echo.Group, peripheralHandler *handler.Peripheral) {
    api.GET("/stats/:id", peripheralHandler)
}
