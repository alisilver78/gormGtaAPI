package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func PrintMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Printf("Middleware triggerd, URL is: %v", c.Request().URL)
		return next(c)
	}
}
