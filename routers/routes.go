package routers

import (
	"gormGtaAPI/handlers"
	"gormGtaAPI/middleware"

	"github.com/labstack/echo/v4"
)

func Register(EP *echo.Echo) {
	webGroup := EP.Group("", middleware.PrintMiddleware)
	routes(webGroup)
}

func routes(EP *echo.Group) {
	EP.GET("/gta", handlers.GetData)
	EP.POST("/gta", handlers.InsertData)
	EP.PUT("/gta/:id", handlers.UpdateData)
	EP.DELETE("/gta/:id", handlers.DeleteRecord)
	EP.GET("/gta/:id", handlers.GetSingle)
	EP.DELETE("/gta", handlers.DeleteAllRecords)
}
