package cell

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestFillShipPositions(t *testing.T) {

	//todo: not working!!
	t.Run("Should create an array of positions for the ship", func(t *testing.T) {
		positions := fillShipPositions(2, 1, 20, 20)

		shipDraw := new(bytes.Buffer)

		board := make([][]bool, 20)

		for i := 0; i < 20; i++ {
			board[i] = make([]bool, 20)
			for j := 0; j < 20; j++ {
				if isFoundInPositions(positions, i, j) {
					board[i][j] = true
				} else {
					board[i][j] = false
				}
			}
		}

		for _, bools := range board {
			for _, b := range bools {
				if b {
					shipDraw.WriteString("x")
				} else {
					shipDraw.WriteString(".")
				}
			}
		}

		_, _ = fmt.Fprintf(os.Stdout, shipDraw.String())
	})
}

func isFoundInPositions(positions []*pixelPoint, i int, j int) bool {
	for _, position := range positions {
		if position.pixelR == i && position.pixelC == j {
			return true
		}
	}

	return false
}
