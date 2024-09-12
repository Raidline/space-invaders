package window

import (
	"bytes"
	"fmt"
	"os"
	"raidline/space-invaders/pkg/assert"
	"raidline/space-invaders/pkg/colors"
)

const MaxRows = 128
const MaxCols = 128

type PixelPoint struct {
	x int
	y int
}

type Cell struct {
	Color    string
	Position PixelPoint
	Valid    bool
}

type Canvas struct {
	field   [][]Cell
	rows    int
	cols    int
	drawBuf *bytes.Buffer
}

const Wall = "\u2588"

func Make(rows, cols int) *Canvas {
	assert.Assert(rows < MaxRows, "Rows should be no bigger than", MaxRows)
	assert.Assert(cols < MaxCols, "Cols should be no bigger than", MaxCols)

	cells := constructBoard(rows, cols)
	return &Canvas{
		field:   cells,
		rows:    rows,
		cols:    cols,
		drawBuf: new(bytes.Buffer),
	}
}

func (c *Canvas) Clear() {
	c.field = constructBoard(c.rows, c.cols)
	c.Draw()
}

func (c *Canvas) Draw() {
	c.drawBuf.Reset()
	// Clear the screen
	_, _ = fmt.Fprint(os.Stdout, "\033[2J")
	// Move the cursor to the home position
	_, _ = fmt.Fprint(os.Stdout, "\033[H")
	c.printBoard()
	c.printStats()
}

func (c *Canvas) printBoard() {
	for _, cells := range c.field {
		for _, cell := range cells {
			if !cell.Valid {
				c.drawBuf.WriteString(Wall)
			} else {
				c.drawBuf.WriteString(" ")
			}
		}
		c.drawBuf.WriteString("\n")
	}
	_, err := fmt.Fprintf(os.Stdout, c.drawBuf.String())
	assert.Assert(err == nil, "We should be able to print the game onto screen")
}

func (c *Canvas) printStats() {

}

func constructBoard(rows, cols int) [][]Cell {
	cells := make([][]Cell, rows)
	for i := 0; i < rows; i++ {
		cells[i] = make([]Cell, cols)
		for j := 0; j < cols; j++ {
			if j == 0 || i == 0 || j == cols-1 || i == rows-1 {
				// this is a border
				cells[i][j] = Cell{
					Color: colors.DarkGray,
					Position: PixelPoint{
						x: i,
						y: j,
					},
					Valid: false,
				}
			} else {
				// we don't support special blocks, so we assume everything is valid for now
				cells[i][j] = Cell{
					Color: colors.Black,
					Position: PixelPoint{
						x: i,
						y: j,
					},
					Valid: true,
				}
			}

		}
	}

	return cells
}
