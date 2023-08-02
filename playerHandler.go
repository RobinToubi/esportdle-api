package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
)

var playerToGuess = Player{
	ID:           "98767991747728851",
	SummonerName: "Faker",
	FirstName:    "Sanghyeok",
	LastName:     "Lee",
	BirthDate:    "1996-05-07T00:00:00Z",
	Team:         "T1",
	Country:      "KR",
	TeamImageUrl: "https://static.lolesports.com/teams/1631819523085_t1-2021-worlds.png",
	Role:         "mid",
	ImageUrl:     "https://static.lolesports.com/players/1655457397135_T1_Faker_784x621.png",
	League:       "LCK",
}

func GuessPlayer(c echo.Context) error {
	var player Player
	var result GuessResult
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
	err = collection.FindOne(currentContext, bson.D{{"id", playerId}}, &opt).Decode(&player)
	if err != nil {
		_ = c.JSON(http.StatusNotFound, GuessResult{})
		return err
	}
	response := Compare(player, playerToGuess)
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
	opt := options.FindOptions{Sort: bson.D{{"summonerName", 1}}}
	cursor, err := collection.Find(currentContext, bson.D{}, &opt)
	if err != nil {
		panic(err)
	}
	if err = cursor.All(currentContext, &players); err != nil {
		return err
	}
	_ = c.JSON(http.StatusOK, &players)
	return nil
}
