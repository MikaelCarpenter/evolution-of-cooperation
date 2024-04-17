package prisonersdilemma_test

import (
	"reflect"
	"testing"

	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/prisonersDilemma"
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

func CheckResults(t *testing.T, result, expectedResult StrategyResult) {
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("GetRoundScore(...) = {Decision:%s PointsAwarded:%d}, want {Decision:%s PointsAwarded:%d}",
			result.Decision, result.PointsAwarded, expectedResult.Decision, expectedResult.PointsAwarded)
	}
}

func TestGetRoundScore(t *testing.T) {

	t.Run("it gives 3pts each if both strats COOPERATE", func(t *testing.T) {
		firstExpectedResult := StrategyResult{
			Decision:      COOPERATE,
			PointsAwarded: 3,
		}
		secondExpectedResult := StrategyResult{
			Decision:      COOPERATE,
			PointsAwarded: 3,
		}

		firstResult, secondResult := GetRoundScore(COOPERATE, COOPERATE)

		CheckResults(t, firstResult, firstExpectedResult)
		CheckResults(t, secondResult, secondExpectedResult)
	})

	t.Run("it gives 1pt each if both strats DEFECT", func(t *testing.T) {
		firstExpectedResult := StrategyResult{
			Decision:      DEFECT,
			PointsAwarded: 1,
		}
		secondExpectedResult := StrategyResult{
			Decision:      DEFECT,
			PointsAwarded: 1,
		}

		firstResult, secondResult := GetRoundScore(DEFECT, DEFECT)

		CheckResults(t, firstResult, firstExpectedResult)
		CheckResults(t, secondResult, secondExpectedResult)
	})

	t.Run("it gives 5pts to first strat if first strat DEFECTS and second strat COOPERATES", func(t *testing.T) {
		firstExpectedResult := StrategyResult{
			Decision:      DEFECT,
			PointsAwarded: 5,
		}
		secondExpectedResult := StrategyResult{
			Decision:      COOPERATE,
			PointsAwarded: 0,
		}

		firstResult, secondResult := GetRoundScore(DEFECT, COOPERATE)

		CheckResults(t, firstResult, firstExpectedResult)
		CheckResults(t, secondResult, secondExpectedResult)
	})

	t.Run("it gives 5pts to second strat if first strat COOPERATES and second strat DEFECTS", func(t *testing.T) {
		firstExpectedResult := StrategyResult{
			Decision:      COOPERATE,
			PointsAwarded: 0,
		}
		secondExpectedResult := StrategyResult{
			Decision:      DEFECT,
			PointsAwarded: 5,
		}

		firstResult, secondResult := GetRoundScore(COOPERATE, DEFECT)

		CheckResults(t, firstResult, firstExpectedResult)
		CheckResults(t, secondResult, secondExpectedResult)
	})
}
