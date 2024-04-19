package main

import (
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/prisonersDilemma"
)

func main() {
	RunTournament(SaveGameResults)
}

func RunTournament(gameRecorder TournamentGameRecorder) {
	keys := make([]string, 0, len(AvailableStrategyMapping))
	for key := range AvailableStrategyMapping {
		keys = append(keys, key)
	}

	for i, key1 := range keys {
		for _, key2 := range keys[i+1:] {
			output := "output/tournament/" + key1 + "_vs_" + key2 + ".json"
			gameRecorder(key1, key2, output, 200)
		}
	}
}

type TournamentGameRecorder func(strat1, strat2, output string, rounds int)
