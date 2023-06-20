package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error when loading .env file")
	}
	server := echo.New()
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world.")
	})
	server.GET("/players", GetPlayers)
	server.POST("/:playerId/guess", GuessPlayer)
	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
