package strategies_test

import (
	"testing"

	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/strategies"
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/testUtils"
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

func TestShameOnMeStrategyFunction(t *testing.T) {
	t.Run("should COOPERATE", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				SHAME_ON_ME: 0,
				MOCK:        0,
			},
			RoundHistory: []RoundResult{},
		}

		AssertDecision(t, ShameOnMeStrategyFunction, input, COOPERATE)
	})

	t.Run("should COOPERATE even if opponent just defected", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				SHAME_ON_ME: 0,
				MOCK:        5,
			},
			RoundHistory: []RoundResult{
				{
					SHAME_ON_ME: StrategyResult{
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

		AssertDecision(t, ShameOnMeStrategyFunction, input, COOPERATE)
	})
}
