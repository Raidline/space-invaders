package window

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"raidline/space-invaders/pkg/assert"
	"raidline/space-invaders/pkg/colors"
	"raidline/space-invaders/pkg/logger"
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

const wall = "\u2588"
const aircraft = "|\n   --=O=--"

func Make(rows, cols int) *Canvas {
	tr, tc := getTerminalSize()

	assert.Assert(rows < tr, "Rows should be no bigger than %d, current terminal row size is %d", tr)
	assert.Assert(cols < tc, "Cols should be no bigger than %d, current terminal cols size is %d", tc)

	cells := constructBoard(rows, cols)
	return &Canvas{
		field:   cells,
		rows:    rows,
		cols:    cols,
		drawBuf: new(bytes.Buffer),
	}
}

func getTerminalSize() (int, int) {
	// Default terminal size
	rows, cols := -1, -1

	// Use the stty command to get the terminal size
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		logger.Warn("Could not open output to cmd")
	}

	_, scanErr := fmt.Sscanf(string(out), "%d %d", &rows, &cols)

	if scanErr != nil {
		logger.Warn("Could not get the terminal size, will continue with default values")
	}

	if rows == -1 || cols == -1 {
		return MaxRows, MaxCols
	}

	return rows, cols
}

func (c *Canvas) Clear() {
	c.field = constructBoard(c.rows, c.cols)
	c.Draw(0)
}

// todo: not happy with passing it like this... i would like to receive a struct with all the engine stats..
// todo: the problem is that the engine needs to call the canvas to draw.. we need to split the circular calling
func (c *Canvas) Draw(fps float64) {
	c.drawBuf.Reset()
	// flush the terminal and go to the top
	_, err := fmt.Fprint(os.Stdout, "\x1b[H\x1b[J")
	if err != nil {
		logger.Error(err.Error())
	}
	c.printBoard()
	c.drawStats(fps)

	_, printErr := fmt.Fprintf(os.Stdout, c.drawBuf.String())
	assert.Assert(printErr == nil, "We should be able to print the game onto screen")
}

func (c *Canvas) printBoard() {
	for _, cells := range c.field {
		for _, cell := range cells {
			if !cell.Valid {
				c.drawBuf.WriteString(wall)
			} else {
				c.drawBuf.WriteString(" ")
			}
		}
		c.drawBuf.WriteString("\n")
	}
}

func (c *Canvas) drawStats(fps float64) {
	_, _ = c.drawBuf.WriteString("--FPS\n")
	_, err := c.drawBuf.WriteString(fmt.Sprintf("FPS: %.2f", fps))
	if err != nil {
		logger.Error("Error printing engine stats %s", err.Error())
	}
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
