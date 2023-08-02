package main

import (
	"github.com/RobinToubi/countries"
	"time"
)

var countryService *countries.CountryData

func init() {
	countryService = countries.InstanceCountry()
}

func Compare(guessedPlayer, currentResult Player) CompareResult {
	return CompareResult{
		Id:      guessedPlayer.ID,
		Role:    guessedPlayer.Role == currentResult.Role,
		Country: IsCountryValidated(guessedPlayer.Country, currentResult.Country),
		League:  guessedPlayer.League == currentResult.League,
		Team:    guessedPlayer.Team == currentResult.Team,
		Older:   isCurrentResultOlder(currentResult.BirthDate, guessedPlayer.BirthDate),
	}
}

func IsCountryValidated(guessedCountry, currentCountry string) string {
	if guessedCountry == currentCountry {
		return "VALID"
	}
	guessedContinent, err := countryService.GetContinentByCountry(guessedCountry)
	if err != nil {
		return ""
	}
	currentContinent, err := countryService.GetContinentByCountry(currentCountry)
	if err != nil {
		return ""
	}
	if currentContinent == guessedContinent {
		return "NEAR"
	}
	return "FAR"
}

func isCurrentResultOlder(currentPlayerDate, guessedPlayerDate string) int {
	currentPlayerTime, _ := time.Parse(time.RFC3339, currentPlayerDate)
	guessedPlayerTime, _ := time.Parse(time.RFC3339, guessedPlayerDate)
	return currentPlayerTime.Compare(guessedPlayerTime)
}
