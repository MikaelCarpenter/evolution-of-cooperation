package strategies_test

import (
	"testing"

	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/strategies"
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/testUtils"
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

func TestFriedmanStrategyFunction(t *testing.T) {
	t.Run("should COOPERATE in the first round", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				FRIEDMAN: 0,
				MOCK:     0,
			},
			RoundHistory: []RoundResult{},
		}

		AssertDecision(t, FriedmanStrategyFunction, input, COOPERATE)
	})

	t.Run("should COOPERATE if opponent has continued to COOPERATE", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				FRIEDMAN: 9,
				MOCK:     9,
			},
			RoundHistory: []RoundResult{
				{
					FRIEDMAN: StrategyResult{
						Decision:      COOPERATE,
						PointsAwarded: 3,
					},
					MOCK: StrategyResult{
						Decision:      COOPERATE,
						PointsAwarded: 3,
					},
				},
				{
					FRIEDMAN: StrategyResult{
						Decision:      COOPERATE,
						PointsAwarded: 3,
					},
					MOCK: StrategyResult{
						Decision:      COOPERATE,
						PointsAwarded: 3,
					},
				},
				{
					FRIEDMAN: StrategyResult{
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

		AssertDecision(t, FriedmanStrategyFunction, input, COOPERATE)
	})

	t.Run("should DEFECT if opponent just chose to DEFECT", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				FRIEDMAN: 0,
				MOCK:     5,
			},
			RoundHistory: []RoundResult{
				{
					FRIEDMAN: StrategyResult{
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

		AssertDecision(t, FriedmanStrategyFunction, input, DEFECT)
	})

	t.Run("should continue to DEFECT even if opponent has begun to cooperate", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				FRIEDMAN: 5,
				MOCK:     5,
			},
			RoundHistory: []RoundResult{
				{
					FRIEDMAN: StrategyResult{
						Decision:      COOPERATE,
						PointsAwarded: 0,
					},
					MOCK: StrategyResult{
						Decision:      DEFECT,
						PointsAwarded: 5,
					},
				},
				{
					FRIEDMAN: StrategyResult{
						Decision:      DEFECT,
						PointsAwarded: 5,
					},
					MOCK: StrategyResult{
						Decision:      COOPERATE,
						PointsAwarded: 0,
					},
				},
			},
		}

		AssertDecision(t, FriedmanStrategyFunction, input, DEFECT)
	})
}
