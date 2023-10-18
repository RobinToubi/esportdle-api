package main

type CompareResult struct {
	Id      string `json:"id"`
	Role    bool   `json:"role"`
	Country string `json:"country"`
	League  bool   `json:"league"`
	Team    bool   `json:"team"`
	Older   int    `json:"older"`
}

type GuessResult struct {
	Id          int           `json:"id"`
	Player      Player        `json:"player"`
	CompareData CompareResult `json:"compare"`
	ValidPlayer Player        `json:"validPlayer,omitempty"`
}
