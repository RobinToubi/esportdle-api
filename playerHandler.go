package main

import (
	"context"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dailyPlayer *Player

func GuessPlayer(c echo.Context) error {
	var player Player
	var result GuessResult
	if dailyPlayer == nil {
		dailyPlayerToGuess, err := getDailyPlayer()
		if err != nil {
			panic(err)
		}
		dailyPlayer = dailyPlayerToGuess
	}
	playerId := c.Param("playerId")
	currentContext := context.TODO()
	client, err := GetConnection(currentContext)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(currentContext); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("lolplayers").Collection("players")
	opt := options.FindOneOptions{}
	err = collection.FindOne(currentContext, bson.D{{Key: "id", Value: playerId}}, &opt).Decode(&player)
	if err != nil {
		_ = c.JSON(http.StatusNotFound, GuessResult{})
		return err
	}
	response := Compare(player, *dailyPlayer)
	id, err := strconv.Atoi(playerId)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, GuessResult{})
		return err
	}
	result = GuessResult{
		Id:          id,
		Player:      player,
		CompareData: response,
	}
	if err != nil {
		panic(err)
	}
	if err = c.JSON(http.StatusOK, result); err != nil {
		return err
	}
	return nil
}

func GetPlayers(c echo.Context) error {
	var players []Player
	players, err := getPlayersArray()
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, nil)
	}
	_ = c.JSON(http.StatusOK, &players)
	return nil
}

func RefreshPlayerToGuess(c echo.Context) error {
	randomPlayer, err := DefineRandomPlayer()
	if err != nil {
		return err
	}
	dailyPlayer = randomPlayer
	return nil
}

func getDailyPlayer() (*Player, error) {
	if dailyPlayer != nil {
		return dailyPlayer, nil
	}
	randomPlayer, err := DefineRandomPlayer()
	if err != nil {
		return nil, err
	}
	return randomPlayer, nil
}

func DefineRandomPlayer() (*Player, error) {
	var players []Player
	players, err := getPlayersArray()
	if err != nil {
		return nil, err
	}
	playerIndex := getRandomNumber(int64(len(players)))
	return &players[playerIndex], nil
}

func getRandomNumber(collectionLength int64) int64 {
	t := time.Now()
	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	random := rand.NewSource(t1.Unix())
	return random.Int63() % collectionLength
}

func getPlayersArray() ([]Player, error) {
	var players []Player
	currentContext := context.TODO()
	client, err := GetConnection(currentContext)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = client.Disconnect(currentContext); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("lolplayers").Collection("players")
	opt := options.FindOptions{Sort: bson.D{{Key: "summonerName", Value: 1}}}
	cursor, err := collection.Find(currentContext, bson.D{}, &opt)
	if err != nil {
		panic(err)
	}
	if err = cursor.All(currentContext, &players); err != nil {
		return nil, err
	}
	return players, nil
}
