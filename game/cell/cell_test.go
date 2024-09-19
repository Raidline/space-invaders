package cell

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFillShipPositions(t *testing.T) {

	//todo: not working!!
	t.Run("Should create an array of positions for the ship", func(t *testing.T) {
		rows := 10
		cols := 10
		positions := fillShipPositions(2, 1, rows, cols)

		shipDraw := new(bytes.Buffer)

		board := make([][]bool, rows)

		wantedStr := `..........
..........
..........
..........
....x.....
...xxx....
..xx.xx...
..........
..........
..........
`

		for i := 0; i < rows; i++ {
			board[i] = make([]bool, cols)
			for j := 0; j < cols; j++ {
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
			shipDraw.WriteString("\n")
		}

		assert.Equal(t, wantedStr, shipDraw.String())
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
