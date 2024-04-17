package main

import (
	"flag"
	"log"

	prisonersdilemma "github.com/MikaelCarpenter/evolution-of-cooperation/internal/prisonersDilemma"
)

var (
	strat1 = flag.String("strat1", "", "The first strategy")
	strat2 = flag.String("strat2", "", "The second strategy")
	rounds = flag.Int("rounds", 1, "The number of rounds")
	output = flag.String("output", "", "The output file path")
)

func main() {
	flag.Parse()

	log.Printf("Running Prisoner's Dilemma: %s vs %s (%d rounds)", *strat1, *strat2, *rounds)
	prisonersdilemma.SaveGameResults(*strat1, *strat2, *output, *rounds)
	log.Printf("Results saved to %s", *output)
}
