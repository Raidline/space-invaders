package window

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"raidline/space-invaders/game"
	"raidline/space-invaders/pkg/assert"
	"raidline/space-invaders/pkg/logger"
	"strconv"
)

type Canvas struct {
	drawBuf *bytes.Buffer
	row     uint16
	col     uint16
}

const wall = "\u2588"
const aircraft = "/\\\t\n      /  \\\n\t / /\\ \\ \t\n\t/_/  \\_\\"

func Make(row, col uint16) *Canvas {
	// todo: we need to listen for the control+c so we can reset the behaviour of cursor showing
	return &Canvas{
		drawBuf: new(bytes.Buffer),
		row:     row,
		col:     col,
	}
}

func (c *Canvas) Flush() {
	_, err := fmt.Fprint(os.Stdout, "\x1b[H\x1b[2J")
	if err != nil {
		logger.Error(err.Error())
	}
}

func (c *Canvas) Clear() {
	c.Flush()
	c.drawBuf.Reset()
}

func (c *Canvas) Draw(fps float64, board [][]game.BoardPoint) {
	// flush the terminal and go to the top
	c.drawBuf.WriteString("\x1b[?25l")
	c.printBoard(board)
	c.drawStats(fps)

	_, printErr := io.Copy(os.Stdout, c.drawBuf)
	c.drawBuf.Reset()
	assert.Assert(printErr == nil, "We should be able to print the game onto screen")
}

var intbuf = make([]byte, 0, 16)

func (c *Canvas) writeCursor(r, col int) {
	c.drawBuf.WriteString("\033[")
	c.drawBuf.Write(strconv.AppendUint(intbuf, uint64(r+1), 10))
	c.drawBuf.WriteString(";")
	c.drawBuf.Write(strconv.AppendUint(intbuf, uint64(col+1), 10))
	c.drawBuf.WriteString("H")
}

func (c *Canvas) printBoard(field [][]game.BoardPoint) {
	for r, cells := range field {
		for col, b := range cells {
			c.writeCursor(r, col)
			if !b.Valid {
				c.drawBuf.WriteString(wall)
			} else {
				c.drawBuf.WriteString(" ")
			}
		}
	}
}

func (c *Canvas) drawStats(fps float64) {
	fpsR := int(c.row)
	c.writeCursor(fpsR-2, 0) // this is basically the line of the last cell column (because it's exclusive values)
	_, _ = c.drawBuf.WriteString("--FPS")
	c.writeCursor(fpsR-1, 0)
	_, err := c.drawBuf.WriteString(fmt.Sprintf("FPS: %.2f", fps))
	if err != nil {
		logger.Error("Error printing engine stats %s", err.Error())
	}
}
