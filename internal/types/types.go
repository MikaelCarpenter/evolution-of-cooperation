package types

/*
Decision represents the output of a strategy.
It can be either COOPERATE (0) or DEFECT (1).
*/
type Decision int

const (
	COOPERATE Decision = iota
	DEFECT
)

func (d Decision) String() string {
	switch d {
	case COOPERATE:
		return "COOPERATE"
	case DEFECT:
		return "DEFECT"
	default:
		return "UNKNOWN"
	}
}

type StrategyResult struct {
	Decision      Decision
	PointsAwarded int
}

/*
StrategyIdentifier is an enum to track the different strategies.
Current strategies include: BASE_LINE, TIT_FOR_TAT
*/
type StrategyIdentifier int

const (
	BASE_LINE StrategyIdentifier = iota
	TIT_FOR_TAT
	MOCK
)

type RoundResult map[StrategyIdentifier]StrategyResult

type GameInput struct {
	Strat1 Strategy
	Strat2 Strategy
	Rounds int
}

/*
# Example

	score := GameScoreStatus{
	    BASE_LINE: 5,
	    TIT_FOR_TAT: 3,
	}
*/
type GameScoreStatus map[StrategyIdentifier]int

/*
Includes the current game score and a history of round results.
The most recent round is the last element in the slice

# Example:

	input := GameState{
			CurrentGameScore: GameScoreStatus{
				BASE_LINE:   0,
				TIT_FOR_TAT: 0,
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
*/
type GameState struct {
	CurrentGameScore GameScoreStatus
	RoundHistory     []RoundResult
}
type Strategy struct {
	Identifier       StrategyIdentifier
	StrategyFunction func(GameState) Decision
}
