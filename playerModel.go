package main

type Player struct {
	ID           string `json:"id" bson:"id"`
	SummonerName string `json:"summonerName" bson:"summonerName"`
	FirstName    string `json:"firstName" bson:"firstName"`
	LastName     string `json:"lastName" bson:"lastName"`
	BirthDate    string `json:"birthDate" bson:"birthDate"`
	Team         string `json:"team" bson:"team"`
	Country      string `json:"country" bson:"country"`
	TeamImageUrl string `json:"teamUrl" bson:"teamImageUrl"`
	Role         string `json:"role" bson:"role"`
	ImageUrl     string `json:"imageUrl" bson:"imageUrl"`
	League       string `json:"league" bson:"league"`
}
