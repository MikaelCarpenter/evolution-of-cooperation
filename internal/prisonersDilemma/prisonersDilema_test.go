package prisonersdilemma_test

import (
	"reflect"
	"testing"

	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/prisonersDilemma"
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/strategies"
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

func TestRunGame(t *testing.T) {
	t.Run("if you pass the same strategy twice, it should return an error", func(t *testing.T) {
		var expected error = nil
		_, received := RunGame(GameInput{Strat1: BaseLineStrategy, Strat2: BaseLineStrategy})

		if expected == received {
			t.Errorf("Expected %v, received %v", expected, received)
		}
	})

	t.Run("if you pass two different strategies, it should return the game score", func(t *testing.T) {
		expected := GameScoreStatus{
			BASE_LINE:   5,
			TIT_FOR_TAT: 0,
		}
		received, _ := RunGame(GameInput{Strat1: BaseLineStrategy, Strat2: TitForTatStrategy})

		if !reflect.DeepEqual(expected, received.CurrentGameScore) {
			t.Errorf("Expected %v, received %v", expected, received.CurrentGameScore)
		}
	})

	t.Run("it should return the results for all the rounds you ask it to run", func(t *testing.T) {
		expected := []RoundResult{
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
		}

		output, _ := RunGame(GameInput{
			Strat1: BaseLineStrategy,
			Strat2: TitForTatStrategy,
			Rounds: 3,
		})
		received := output.RoundHistory

		for i := range expected {
			if i >= len(received) {
				t.Errorf("Expected at least %d rounds, but got %d rounds", len(expected), len(received))
				break
			}
			if !reflect.DeepEqual(expected[i], received[i]) {
				t.Errorf("Mismatch in round %d (index %d): expected %v, but got %v", i+1, i, expected[i], received[i])
			}
		}
	})
}
