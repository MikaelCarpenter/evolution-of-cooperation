package strategies_test

import (
	"testing"

	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/strategies"
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/testUtils"
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

func TestCopyCatStrategyFunction(t *testing.T) {
	t.Run("should COOPERATE in the first round", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				COPY_CAT: 0,
				MOCK:     0,
			},
			RoundHistory: []RoundResult{},
		}

		AssertDecision(t, CopyCatStrategyFunction, input, COOPERATE)
	})

	t.Run("should COOPERATE if opponent COOPERATE-ed in the last round", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				COPY_CAT: 3,
				MOCK:     3,
			},
			RoundHistory: []RoundResult{
				{
					COPY_CAT: StrategyResult{
						Decision:      COOPERATE,
						PointsAwarded: 3,
					},
					MOCK: StrategyResult{
						Decision:      COOPERATE,
						PointsAwarded: 3,
					},
				},
			},
		}

		AssertDecision(t, CopyCatStrategyFunction, input, COOPERATE)
	})

	t.Run("should DEFECT if opponent DEFECT-ed in the last round", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				COPY_CAT: 0,
				MOCK:     5,
			},
			RoundHistory: []RoundResult{
				{
					COPY_CAT: StrategyResult{
						Decision:      COOPERATE,
						PointsAwarded: 0,
					},
					MOCK: StrategyResult{
						Decision:      DEFECT,
						PointsAwarded: 5,
					},
				},
			},
		}

		AssertDecision(t, CopyCatStrategyFunction, input, DEFECT)
	})
}
