package testutils

import (
	"testing"

	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

func AssertDecision(t *testing.T, strategy StrategyFunction, input GameState, expected Decision) {
	received := strategy(input)

	if expected != received {
		t.Errorf("Expected %s, received %s", expected, received)
	}
}
