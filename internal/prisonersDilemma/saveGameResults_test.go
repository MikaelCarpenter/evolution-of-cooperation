package prisonersdilemma_test

import (
	"os"
	"testing"

	. "github.com/MikaelCarpenter/evolution-of-cooperation/internal/prisonersDilemma"
)

func TestSaveGameResults(t *testing.T) {
	t.Run("it should run successfully", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "gameState")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer os.Remove(tempFile.Name())

		SaveGameResults("BaseLineStrategy", "TitForTatStrategy", tempFile.Name(), 1)

		content, err := os.ReadFile(tempFile.Name())
		if err != nil {
			t.Fatalf("Failed to read temporary file: %v", err)
		}

		if string(content) != expectedOutput {
			t.Errorf("Unexpected output: got %v, want %v", string(content), expectedOutput)
		}
	})
}

const expectedOutput = `{
  "CurrentGameScore": {
    "0": 5,
    "1": 0
  },
  "RoundHistory": [
    {
      "0": {
        "Decision": 1,
        "PointsAwarded": 5
      },
      "1": {
        "Decision": 0,
        "PointsAwarded": 0
      }
    }
  ]
}`
