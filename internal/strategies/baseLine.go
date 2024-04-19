package strategies

import (
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

var BaseLineStrategy = Strategy{
	Identifier:       BASE_LINE,
	StrategyFunction: BaseLineStrategyFunction,
}

func BaseLineStrategyFunction(input GameState) Decision {
	return DEFECT
}
