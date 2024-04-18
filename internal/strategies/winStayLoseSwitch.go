package strategies

import (
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

var WinStayLoseSwitchStrategy = Strategy{
	Identifier:       WIN_STAY_LOSE_SWITCH,
	StrategyFunction: WinStayLoseSwitchStrategyFunction,
}

func WinStayLoseSwitchStrategyFunction(input GameState) Decision {
	return COOPERATE
}
