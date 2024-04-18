package strategies

import (
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

var FriedmanStrategy = Strategy{
	Identifier:       FRIEDMAN,
	StrategyFunction: FriedmanStrategyFunction,
}

func FriedmanStrategyFunction(input GameState) Decision {
	if len(input.RoundHistory) == 0 {
		return COOPERATE
	}

	var opponentIdentifier StrategyIdentifier
	for identifier := range input.CurrentGameScore {
		if identifier != FRIEDMAN {
			opponentIdentifier = identifier
			break
		}
	}

	for _, roundResult := range input.RoundHistory {
		if roundResult[opponentIdentifier].Decision == DEFECT {
			return DEFECT
		}
	}

	return COOPERATE
}
