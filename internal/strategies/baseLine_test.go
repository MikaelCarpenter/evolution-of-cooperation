package strategies_test

import (
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/strategies"
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/testUtils"
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"

	"testing"
)

func TestBaseLineStrategy(t *testing.T) {
	t.Run("defects if opponent cooperated last round", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				BASE_LINE:   5,
				TIT_FOR_TAT: 0,
			},
			RoundHistory: []RoundResult{
				{
					BASE_LINE: StrategyResult{
						Decision:      DEFECT,
						PointsAwarded: 5,
					},
					TIT_FOR_TAT: StrategyResult{
						Decision:      COOPERATE,
						PointsAwarded: 0,
					},
				},
			},
		}

		AssertDecision(t, BaseLineStrategyFunction, input, DEFECT)
	})

	t.Run("defects if opponent defected last round", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				BASE_LINE:   1,
				TIT_FOR_TAT: 1,
			},
			RoundHistory: []RoundResult{
				{
					BASE_LINE: StrategyResult{
						Decision:      DEFECT,
						PointsAwarded: 1,
					},
					TIT_FOR_TAT: StrategyResult{
						Decision:      DEFECT,
						PointsAwarded: 1,
					},
				},
			},
		}

		AssertDecision(t, BaseLineStrategyFunction, input, DEFECT)
	})
}
