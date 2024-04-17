package prisonersdilemma

import (
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

func GetRoundScore(decisionOne, decisionTwo Decision) (firstResult, secondResult StrategyResult) {
	if decisionOne == COOPERATE {
		if decisionTwo == COOPERATE {
			firstResult = StrategyResult{
				Decision:      COOPERATE,
				PointsAwarded: 3,
			}
			secondResult = StrategyResult{
				Decision:      COOPERATE,
				PointsAwarded: 3,
			}
		} else {
			firstResult = StrategyResult{
				Decision:      COOPERATE,
				PointsAwarded: 0,
			}
			secondResult = StrategyResult{
				Decision:      DEFECT,
				PointsAwarded: 5,
			}
		}
	} else {
		if decisionTwo == COOPERATE {
			firstResult = StrategyResult{
				Decision:      DEFECT,
				PointsAwarded: 5,
			}
			secondResult = StrategyResult{
				Decision:      COOPERATE,
				PointsAwarded: 0,
			}
		} else {
			firstResult = StrategyResult{
				Decision:      DEFECT,
				PointsAwarded: 1,
			}
			secondResult = StrategyResult{
				Decision:      DEFECT,
				PointsAwarded: 1,
			}
		}
	}

	return
}
