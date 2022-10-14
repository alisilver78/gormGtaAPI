package main

import (
	"gormGtaAPI/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
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

	e.Pre(middleware.RemoveTrailingSlash())
	routers.Register(e)

	e.Logger.Fatal(e.Start(":8080"))

}
