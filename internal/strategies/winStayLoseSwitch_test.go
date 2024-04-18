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
				WIN_STAY_LOSE_SWITCH: 0,
				MOCK:                 0,
			},
			RoundHistory: []RoundResult{},
		}

		AssertDecision(t, WinStayLoseSwitchStrategyFunction, input, COOPERATE)
	})

	t.Run("should stick with it's decision if it got same or more points than opponent", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				WIN_STAY_LOSE_SWITCH: 3,
				MOCK:                 3,
			},
			RoundHistory: []RoundResult{
				{
					WIN_STAY_LOSE_SWITCH: StrategyResult{
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

		AssertDecision(t, WinStayLoseSwitchStrategyFunction, input, COOPERATE)
	})

	t.Run("should switch decisions if it got less points than opponent", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				WIN_STAY_LOSE_SWITCH: 0,
				MOCK:                 5,
			},
			RoundHistory: []RoundResult{
				{
					WIN_STAY_LOSE_SWITCH: StrategyResult{
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

		AssertDecision(t, WinStayLoseSwitchStrategyFunction, input, DEFECT)
	})
}
