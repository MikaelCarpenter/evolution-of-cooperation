package prisonersdilemma

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/strategies"
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

func SaveGameResults(strat1, strat2, output string, rounds int) {
	flag.Parse()

	if strat1 == "" || strat2 == "" || output == "" {
		fmt.Println("You must provide two strategies and an output file path")
		os.Exit(1)
	}

	strategyMapping := map[string]Strategy{
		"BaseLineStrategy":          BaseLineStrategy,
		"TitForTatStrategy":         TitForTatStrategy,
		"MaslowStrategy":            MaslowStrategy,
		"WinStayLoseSwitchStrategy": WinStayLoseSwitchStrategy,
	}

	gameInput := GameInput{
		Strat1: strategyMapping[strat1],
		Strat2: strategyMapping[strat2],
		Rounds: rounds,
	}

	gameState, err := RunGame(gameInput)
	if err != nil {
		log.Fatalf("Failed to run game: %v", err)
	}

	gameStateJson, err := json.MarshalIndent(gameState, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal game state: %v", err)
	}

	err = os.WriteFile(output, gameStateJson, 0644)
	if err != nil {
		log.Fatalf("Failed to write game state to file: %v", err)
	}
}
