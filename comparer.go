package main

import (
	"time"

	"github.com/RobinToubi/countries"
)

var countryService *countries.CountryData

func init() {
	countryService = countries.InstanceCountry()
}

// Comparing the guessedPlayer to the result stored in API.
func Compare(guessedPlayer, currentResult Player) CompareResult {
	return CompareResult{
		Id:      guessedPlayer.ID,
		Role:    guessedPlayer.Role == currentResult.Role,
		Country: IsCountryValidated(guessedPlayer.Country, currentResult.Country),
		League:  guessedPlayer.League == currentResult.League,
		Team:    guessedPlayer.Team == currentResult.Team,
		Older:   isCurrentResultOlder(guessedPlayer.BirthDate, currentResult.BirthDate),
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

func isCurrentResultOlder(guessedPlayerDate, currentPlayerDate string) int {
	currentPlayerTime, _ := time.Parse(time.RFC3339, currentPlayerDate)
	guessedPlayerTime, _ := time.Parse(time.RFC3339, guessedPlayerDate)
	currentPlayerAge := getAge(currentPlayerTime, time.Now())
	guessedPlayerAge := getAge(guessedPlayerTime, time.Now())
	if currentPlayerAge == guessedPlayerAge {
		return 0 // The current result has the same age as the guessed result
	}
	if currentPlayerAge < guessedPlayerAge {
		return -1 // The current result is younger than the guessed result
	}
	return 1 // The current result is older than the guessed result
}

func getAge(playerDate time.Time, to time.Time) int {
	return int(to.Sub(playerDate).Hours()) / 24 / 365
}
