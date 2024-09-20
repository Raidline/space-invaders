package cell

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFillPositions(t *testing.T) {

	t.Run("Should create an array of Positions for the ship", func(t *testing.T) {
		rows := 10
		cols := 10
		positions := fillShipPositions(2, rows, cols)

		shipDraw := new(bytes.Buffer)

		board := make([][]bool, rows)

		wantedStr := `..........
..........
..........
..........
..........
..........
....x.....
...xxx....
..xxxxx...
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

	t.Run("Should create an array of Positions for an enemy", func(t *testing.T) {
		cols := 10
		rows := 10

		positions := fillEnemyPositions(3, 5, 1, 2)

		draw := new(bytes.Buffer)

		board := make([][]bool, rows)

		wantedStr := `..........
..xxxxx...
..xxxxx...
..xxxxx...
..........
..........
..........
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
					draw.WriteString("x")
				} else {
					draw.WriteString(".")
				}
			}
			draw.WriteString("\n")
		}

		assert.Equal(t, wantedStr, draw.String())

	})
}

func isFoundInPositions(positions []*PixelPoint, i int, j int) bool {
	for _, position := range positions {
		if position.PixelR == i && position.PixelC == j {
			return true
		}
	}

	return false
}
