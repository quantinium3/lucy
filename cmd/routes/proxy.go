package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/quantinium3/lucy/cmd/handler"
)

func SetupProxyRoutes(api *echo.Group) {
    api.GET("/lastfm", handler.GetLastFMTracks)

    api.GET("/wakatime/currentproject", handler.GetCurrentProject)
    api.GET("/wakatime/os", handler.GetOperatingSystems)
    api.GET("/wakatime/machine", handler.GetMachine)
    api.GET("/wakatime/languages", handler.GetLanguages)
    api.GET("/wakatime/editor", handler.GetEditor)
    api.GET("/wakatime/time", handler.GetTime)

    api.GET("/todoist/tasks", handler.GetTasks)
}
