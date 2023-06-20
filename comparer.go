package main

func Compare(guessedPlayer, currentResult Player) GuessAnswer {
	return GuessAnswer{
		Role:              guessedPlayer.Role == currentResult.Role,
		Country:           "NEAR",
		League:            guessedPlayer.League == currentResult.League,
		WorldsApparitions: "LESS",
	}
}
