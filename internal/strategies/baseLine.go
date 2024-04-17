package strategies

import (
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

var BaseLineStrategy = Strategy{
	Identifier:       BASE_LINE,
	StrategyFunction: BaseLineStrategyFunction,
}

// If you were only playing 1 round, and you had no idea what your
// opponent would do, then DEFECT-ing will get you the most points
// regardless of what your opponent decides.
func BaseLineStrategyFunction(input GameState) Decision {
	return DEFECT
}
