package main

import (
	"testing"
	"time"
)

func TestGetAge(t *testing.T) {
	currentResultBirthDate, _ := time.Parse(time.RFC3339, "2001-09-26T00:00:00Z")
	fakedCurrentDay, _ := time.Parse(time.RFC3339, "2015-08-24T00:00:00Z")
	age := getAge(currentResultBirthDate, fakedCurrentDay)
	if age != 13 {
		t.Errorf("The age is not as expected.\nexpected : 13 | current : %d", age)
	}
}

func TestIsCurrentResultOlder_SameAge(t *testing.T) {
	currentResultBirthDate, guessedBirthDate := "2001-09-26T00:00:00Z", "2002-08-24T00:00:00Z"
	expected := 0
	current := isCurrentResultOlder(guessedBirthDate, currentResultBirthDate)
	if expected != current {
		t.Error("Current result isn't older than the guessed result")
	}
}
