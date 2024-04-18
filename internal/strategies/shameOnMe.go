package strategies

import (
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

var ShameOnMeStrategy = Strategy{
	Identifier:       SHAME_ON_ME,
	StrategyFunction: ShameOnMeStrategyFunction,
}

func ShameOnMeStrategyFunction(input GameState) Decision {
	return COOPERATE
}
