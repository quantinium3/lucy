package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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
	peripheralService := handler.PeripheralService(db)

	e := echo.New()
	api := e.Group("/api/v1")
	routes.SetupUserRoutes(api, userService)
	routes.SetupPeripheralRoutes(api, peripheralService)
	routes.SetupProxyRoutes(api)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
