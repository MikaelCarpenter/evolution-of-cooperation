package strategies_test

import (
	"testing"

	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/strategies"
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

func TestMaslowStrategyFunction(t *testing.T) {
	t.Run("it should COOPERATE in the first round", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				MASLOW:    0,
				BASE_LINE: 0,
			},
			RoundHistory: []RoundResult{},
		}

		expected := COOPERATE
		received := MaslowStrategyFunction(input)

		if expected != received {
			t.Errorf("Expected decision: %v. Received: %v", expected, received)
		}
	})

	t.Run("it should COOPERATE if currently winning", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				MASLOW: 1,
				MOCK:   0,
			},
			RoundHistory: []RoundResult{},
		}

		expected := COOPERATE
		received := MaslowStrategyFunction(input)

		if expected != received {
			t.Errorf("Expected decision: %v. Received: %v", expected, received)
		}
	})

	t.Run("it should COOPERATE if currently tied", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				MASLOW: 3,
				MOCK:   3,
			},
			RoundHistory: []RoundResult{},
		}

		expected := COOPERATE
		received := MaslowStrategyFunction(input)

		if expected != received {
			t.Errorf("Expected decision: %v. Received: %v", expected, received)
		}
	})

	t.Run("it should DEFECT if currently losing", func(t *testing.T) {
		input := GameState{
			CurrentGameScore: GameScoreStatus{
				MASLOW: 0,
				MOCK:   5,
			},
			RoundHistory: []RoundResult{},
		}

		expected := DEFECT
		received := MaslowStrategyFunction(input)

		if expected != received {
			t.Errorf("Expected decision: %v. Received: %v", expected, received)
		}
	})
}
