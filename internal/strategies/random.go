package strategies

import (
	"math/rand"

	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

var RandomStrategy = Strategy{
	Identifier:       RANDOM,
	StrategyFunction: RandomStrategyFunction,
}

func RandomStrategyFunction(input GameState) Decision {
	var rand = rand.New(rand.NewSource(42))
	if rand.Float64() < 0.5 {
		return COOPERATE
	} else {
		return DEFECT
	}
}
