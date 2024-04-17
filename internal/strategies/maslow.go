package strategies

import (
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

var MaslowStrategy = Strategy{
	Identifier:       MASLOW,
	StrategyFunction: MaslowStrategyFunction,
}

func MaslowStrategyFunction(input GameState) Decision {
	myScore := input.CurrentGameScore[MASLOW]

	var opponentsScore int
	for identifier, score := range input.CurrentGameScore {
		if identifier != MASLOW {
			opponentsScore = score
			break
		}
	}

	if myScore >= opponentsScore {
		return COOPERATE
	} else {
		return DEFECT
	}
}
