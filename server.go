package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error when loading .env file")
	}
	server := echo.New()
	server.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world.")
	})
	server.POST("/:playerId/guess", GuessPlayer)
	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", ":8080")))
}
