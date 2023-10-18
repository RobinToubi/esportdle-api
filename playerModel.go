package main

type Player struct {
	ID           string `json:"id,omitempty" bson:"id"`
	SummonerName string `json:"summonerName,omitempty" bson:"summonerName"`
	FirstName    string `json:"firstName,omitempty" bson:"firstName"`
	LastName     string `json:"lastName,omitempty" bson:"lastName"`
	BirthDate    string `json:"birthDate,omitempty" bson:"birthDate"`
	Team         string `json:"team,omitempty" bson:"team"`
	Country      string `json:"country,omitempty" bson:"country"`
	TeamImageUrl string `json:"teamUrl,omitempty" bson:"teamImageUrl"`
	Role         string `json:"role,omitempty" bson:"role"`
	ImageUrl     string `json:"imageUrl,omitempty" bson:"imageUrl"`
	League       string `json:"league,omitempty" bson:"league"`
}
