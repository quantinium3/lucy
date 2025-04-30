package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/quantinium3/lucy/cmd/database"
	"github.com/quantinium3/lucy/cmd/handler"
	"github.com/quantinium3/lucy/cmd/routes"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

    userService := handler.UserService(db)

	e := echo.New()
    api := e.Group("/api/v1")
    routes.SetupUserRoutes(api, userService);
    e.Logger.Fatal(e.Start(":1232"))
}
