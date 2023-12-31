package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func getPlayers() ([]Player, error) {
	var players []Player
	content, err := os.ReadFile("./players.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &players)
	if err != nil {
		panic(err)
	}
	return players, nil
}

func GuessPlayer(c echo.Context) error {
	var guessedPlayer *Player
	var result GuessResult
	players, err := getPlayers()
	if err != nil {
		panic(err)
	}
	validPlayer := getDailyPlayer(players)
	playerId := c.Param("playerId")
	for _, p := range players {
		if p.ID == playerId {
			guessedPlayer = &p
			break
		}
	}
	c.Logger().Info("Guess the player : %s", guessedPlayer.SummonerName)
	compare := Compare(guessedPlayer, validPlayer)
	id, err := strconv.Atoi(playerId)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, GuessResult{})
		return err
	}
	result = GuessResult{
		Id:          id,
		Player:      *guessedPlayer,
		CompareData: compare,
		ValidPlayer: Player{},
	}
	guesses := c.QueryParam("guesses")
	if os.Getenv("MAX_GUESSES") == guesses || validPlayer.ID == guessedPlayer.ID {
		result.ValidPlayer = *validPlayer
	}
	if err = c.JSON(http.StatusOK, result); err != nil {
		return err
	}
	return nil
}

func getDailyPlayer(players []Player) *Player {
	player, err := defineRandomPlayer(players)
	if err != nil {
		panic(err)
	}
	return player
}

func defineRandomPlayer(players []Player) (*Player, error) {
	playerIndex := getRandomNumber(int64(len(players)))
	return &players[playerIndex], nil
}

func getRandomNumber(collectionLength int64) int64 {
	location, _ := time.LoadLocation("Europe/Paris")
	t := time.Now().In(location)
	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, location)
	random := rand.NewSource(t1.Unix())
	return random.Int63() % collectionLength
}
