package strategies

import (
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

var TitForTatStrategy = Strategy{
	Identifier:       TIT_FOR_TAT,
	StrategyFunction: TitForTatStrategyFunction,
}

func TitForTatStrategyFunction(input GameState) Decision {
	roundHistoryLength := len(input.RoundHistory)

	if roundHistoryLength == 0 {
		return COOPERATE
	}

	previousRound := input.RoundHistory[roundHistoryLength-1]

	var opponentPreviousDecision Decision
	for identifier, strategyResult := range previousRound {
		if identifier != TIT_FOR_TAT {
			opponentPreviousDecision = strategyResult.Decision
			break
		}
	}

	if opponentPreviousDecision == DEFECT {
		return DEFECT
	} else {
		return COOPERATE
	}
}
