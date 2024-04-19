package prisonersdilemma

import (
	"errors"

	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/strategies"
	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/types"
)

var AvailableStrategyMapping = map[string]Strategy{
	"BaseLineStrategy":          BaseLineStrategy,
	"TitForTatStrategy":         TitForTatStrategy,
	"MaslowStrategy":            MaslowStrategy,
	"WinStayLoseSwitchStrategy": WinStayLoseSwitchStrategy,
	"ShameOnMeStrategy":         ShameOnMeStrategy,
	"CopyCatStrategy":           CopyCatStrategy,
	"RandomStrategy":            RandomStrategy,
	"FriedmanStrategy":          FriedmanStrategy,
}

func RunGame(input GameInput) (GameState, error) {
	if input.Strat1.Identifier == input.Strat2.Identifier {
		return GameState{}, errors.New("cannot pass two of the same strategies")
	}

	if input.Rounds == 0 {
		input.Rounds = 1
	}

	gameState := GameState{
		CurrentGameScore: GameScoreStatus{
			input.Strat1.Identifier: 0,
			input.Strat2.Identifier: 0,
		},
		RoundHistory: []RoundResult{},
	}

	for i := 0; i < input.Rounds; i++ {
		strat1Decision, strat2Decision := input.Strat1.StrategyFunction(gameState), input.Strat2.StrategyFunction(gameState)
		strat1Result, strat2Result := GetRoundScore(strat1Decision, strat2Decision)

		gameState.CurrentGameScore[input.Strat1.Identifier] += strat1Result.PointsAwarded
		gameState.CurrentGameScore[input.Strat2.Identifier] += strat2Result.PointsAwarded

		roundResult := RoundResult{
			input.Strat1.Identifier: StrategyResult{
				Decision:      strat1Decision,
				PointsAwarded: strat1Result.PointsAwarded,
			},
			input.Strat2.Identifier: StrategyResult{
				Decision:      strat2Decision,
				PointsAwarded: strat2Result.PointsAwarded,
			},
		}

		gameState.RoundHistory = append(gameState.RoundHistory, roundResult)
	}

	return gameState, nil
}
