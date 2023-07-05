package main

import "github.com/RobinToubi/countries"

func Compare(guessedPlayer, currentResult Player) CompareResult {
	return CompareResult{
		Id:      guessedPlayer.ID,
		Role:    guessedPlayer.Role == currentResult.Role,
		Country: IsCountryValidated(guessedPlayer.Country, currentResult.Country),
		League:  guessedPlayer.League == currentResult.League,
	}
}

func IsCountryValidated(guessedCountry, currentCountry string) string {
	if guessedCountry == currentCountry {
		return "VALID"
	}
	countryService := countries.InstanceCountry()
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
