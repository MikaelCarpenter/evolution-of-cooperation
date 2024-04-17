package strategies_test

import (
	"testing"

	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/strategies"
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

func TestTitForTatStrategyFucntion(t *testing.T) {
	t.Run("it should COOPERATE in the first round", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				TIT_FOR_TAT: 0,
				BASE_LINE:   0,
			},
			RoundHistory: []RoundResult{},
		}

		expected := COOPERATE
		received := TitForTatStrategyFunction(input)

		if expected != received {
			t.Errorf("Expected decision: %v. Received: %v", expected, received)
		}
	})

	t.Run("it should DEFECT if the opponent DEFECTs in the round prior", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				TIT_FOR_TAT: 0,
				BASE_LINE:   5,
			},
			RoundHistory: []RoundResult{
				{
					TIT_FOR_TAT: StrategyResult{
						Decision:      COOPERATE,
						PointsAwarded: 0,
					},
					BASE_LINE: StrategyResult{
						Decision:      DEFECT,
						PointsAwarded: 5,
					},
				},
			},
		}

		expected := DEFECT
		received := TitForTatStrategyFunction(input)

		if expected != received {
			t.Errorf("Expected decision: %v. Received: %v", expected, received)
		}
	})

	t.Run("it should continue to DEFECT if the opponent continues to DEFECT", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				TIT_FOR_TAT: 1,
				BASE_LINE:   6,
			},
			RoundHistory: []RoundResult{
				{
					TIT_FOR_TAT: StrategyResult{
						Decision:      COOPERATE,
						PointsAwarded: 0,
					},
					BASE_LINE: StrategyResult{
						Decision:      DEFECT,
						PointsAwarded: 5,
					},
				},
				{
					TIT_FOR_TAT: StrategyResult{
						Decision:      DEFECT,
						PointsAwarded: 1,
					},
					BASE_LINE: StrategyResult{
						Decision:      DEFECT,
						PointsAwarded: 1,
					},
				},
			},
		}

		expected := DEFECT
		received := TitForTatStrategyFunction(input)

		if expected != received {
			t.Errorf("Expected decision: %v. Received: %v", expected, received)
		}
	})

	t.Run("it should return to COOPERATE-ing if the opponent decides to COOPERATE in the previous round", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				TIT_FOR_TAT: 6,
				BASE_LINE:   6,
			},
			RoundHistory: []RoundResult{
				{
					TIT_FOR_TAT: StrategyResult{
						Decision:      COOPERATE,
						PointsAwarded: 0,
					},
					BASE_LINE: StrategyResult{
						Decision:      DEFECT,
						PointsAwarded: 5,
					},
				},
				{
					TIT_FOR_TAT: StrategyResult{
						Decision:      DEFECT,
						PointsAwarded: 1,
					},
					BASE_LINE: StrategyResult{
						Decision:      DEFECT,
						PointsAwarded: 1,
					},
				},
				{
					TIT_FOR_TAT: StrategyResult{
						Decision:      DEFECT,
						PointsAwarded: 5,
					},
					BASE_LINE: StrategyResult{
						Decision:      COOPERATE,
						PointsAwarded: 0,
					},
				},
			},
		}

		expected := COOPERATE
		received := TitForTatStrategyFunction(input)

		if expected != received {
			t.Errorf("Expected decision: %v. Received: %v", expected, received)
		}
	})

}
