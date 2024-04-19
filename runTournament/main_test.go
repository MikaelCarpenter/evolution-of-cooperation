package main

import (
	"testing"

	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/prisonersDilemma"
)

func TestRunTournament(t *testing.T) {
	numOfStrategies := len(AvailableStrategyMapping)
	numUniqueMatchups := numOfStrategies * (numOfStrategies - 1) / 2

	RunTournament(SaveGameResultsSpy)

	if len(SaveGameResultsCalls) != numUniqueMatchups {
		t.Errorf("Expected %d calls to SaveGameResults, got %d", numUniqueMatchups, len(SaveGameResultsCalls))
	}

	for key1 := range AvailableStrategyMapping {
		count := 0
		for _, call := range SaveGameResultsCalls {
			if call.key1 == key1 || call.key2 == key1 {
				count++
			}
		}
		if count != numOfStrategies-1 {
			t.Errorf("Expected %s to play %d games, got %d", key1, numOfStrategies-1, count)
		}
	}
}

var SaveGameResultsCalls []struct {
	key1, key2, output string
	rounds             int
}

func SaveGameResultsSpy(key1, key2, output string, rounds int) {
	SaveGameResultsCalls = append(SaveGameResultsCalls, struct {
		key1, key2, output string
		rounds             int
	}{key1, key2, output, rounds})
}
