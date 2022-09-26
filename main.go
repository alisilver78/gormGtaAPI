package main

import (
	"gormGtaAPI/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/gta", handlers.GetData)
	e.POST("/gta", handlers.InsertData)
	e.PUT("/gta/:id", handlers.UpdateData)
	e.DELETE("/gta/:id", handlers.DeleteRecord)
	e.GET("/gta/:id", handlers.GetSingle)

	e.Logger.Fatal(e.Start(":8080"))

}
