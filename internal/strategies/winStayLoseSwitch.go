package strategies

import (
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

var WinStayLoseSwitchStrategy = Strategy{
	Identifier:       WIN_STAY_LOSE_SWITCH,
	StrategyFunction: WinStayLoseSwitchStrategyFunction,
}

func WinStayLoseSwitchStrategyFunction(input GameState) Decision {
	if len(input.RoundHistory) == 0 {
		return COOPERATE
	}

	lastRound := input.RoundHistory[len(input.RoundHistory)-1]
	prevDecision := lastRound[WIN_STAY_LOSE_SWITCH].Decision

	if lastRound[WIN_STAY_LOSE_SWITCH].PointsAwarded >= lastRound[MOCK].PointsAwarded {
		return prevDecision
	} else if prevDecision == COOPERATE {
		return DEFECT
	} else {
		return COOPERATE
	}
}
