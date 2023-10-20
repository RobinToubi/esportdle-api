package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	_ = godotenv.Load()
	server := echo.New()
	server.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world.")
	})
	server.POST("/:playerId/guess", GuessPlayer)
	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", "12000")))
}
