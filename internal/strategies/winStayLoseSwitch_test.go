package strategies_test

import (
	"testing"

	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/strategies"
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/testUtils"
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

func TestWinStayLoseSwitchStrategyFunction(t *testing.T) {
	t.Run("should cooperate in the first round", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				BASE_LINE:   0,
				TIT_FOR_TAT: 0,
			},
			RoundHistory: []RoundResult{},
		}

		AssertDecision(t, WinStayLoseSwitchStrategyFunction, input, COOPERATE)
	})
}
