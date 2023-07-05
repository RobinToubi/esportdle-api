package main

type CompareResult struct {
	Id      string `json:"id"`
	Role    bool   `json:"role"`
	Country string `json:"country"`
	League  bool   `json:"league"`
}
