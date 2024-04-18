package strategies

import (
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

var CopyCatStrategy = Strategy{
	Identifier:       COPY_CAT,
	StrategyFunction: CopyCatStrategyFunction,
}

func CopyCatStrategyFunction(input GameState) Decision {
	if len(input.RoundHistory) == 0 {
		return COOPERATE
	}

	var opponentLastRoundDecision Decision
	for identifier, roundResult := range input.RoundHistory[len(input.RoundHistory)-1] {
		if identifier != COPY_CAT {
			opponentLastRoundDecision = roundResult.Decision
			break
		}
	}

	if opponentLastRoundDecision == DEFECT {
		return DEFECT
	}

	return COOPERATE
}
