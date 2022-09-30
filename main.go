package main

import (
	"gormGtaAPI/handlers"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("gta.db"), &gorm.Config{})
	if err != nil {
		log.Panicf("Faild to connect database: %v", err)
	}
}

func main() {
	e := echo.New()

	e.GET("/gta", handlers.GetData)
	e.POST("/gta", handlers.InsertData)
	e.PUT("/gta/:id", handlers.UpdateData)
	e.DELETE("/gta/:id", handlers.DeleteRecord)
	e.GET("/gta/:id", handlers.GetSingle)
	e.DELETE("/gta", handlers.DeleteAllRecords)

	e.Logger.Fatal(e.Start(":8080"))

}
