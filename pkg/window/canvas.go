package window

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"raidline/space-invaders/pkg/assert"
	"raidline/space-invaders/pkg/colors"
	"raidline/space-invaders/pkg/logger"
	"strconv"
	"syscall"
	"unsafe"
)

const MaxRows = 128
const MaxCols = 128

type PixelPoint struct {
	pixelR int
	pixelC int
}

type Cell struct {
	Color    string
	Position PixelPoint
	Valid    bool
}

type Canvas struct {
	field   [][]Cell
	drawBuf *bytes.Buffer
	ws      *winsize
}

type winsize struct {
	Row uint16
	Col uint16
}

const wall = "\u2588"
const aircraft = "/\\\t\n      /  \\\n\t / /\\ \\ \t\n\t/_/  \\_\\"

func Make() *Canvas {
	ws := getWinsize()

	//todo: stats takes 2 lines (we should have this calculated)
	cells := constructBoard(int(ws.Row-2), int(ws.Col))

	// todo: we need to listen for the control+c so we can reset the behaviour
	//_, err := fmt.Fprint(os.Stdout, "\x1b[?25l") // hide the cursor
	//
	//// \x1b[?12l\x1b[?25h -> shows the cursor
	//
	//if err != nil {
	//	logger.Warn("Cursor could not be hidden, will be shown")
	//}
	return &Canvas{
		field:   cells,
		ws:      ws,
		drawBuf: new(bytes.Buffer),
	}
}

func getWinsize() *winsize {

	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		logger.Error("Could not get terminal information")
		panic(errno)
	}

	if ws.Row <= 0 {
		ws.Row = MaxRows
	}

	if ws.Col <= 0 {
		ws.Col = MaxCols
	}

	return ws
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

// todo: not happy with passing it like this... i would like to receive a struct with all the engine stats..
// todo: the problem is that the engine needs to call the canvas to draw.. we need to split the circular calling
func (c *Canvas) Draw(fps float64) {
	// flush the terminal and go to the top
	c.drawBuf.WriteString("\x1b[?25l")
	c.printBoard()
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

func (c *Canvas) printBoard() {
	for _, cells := range c.field {
		for _, cell := range cells {
			c.writeCursor(cell.Position.pixelR, cell.Position.pixelC)
			if !cell.Valid {
				c.drawBuf.WriteString(wall)
			} else {
				c.drawBuf.WriteString(" ")
			}
		}
	}
}

func (c *Canvas) drawStats(fps float64) {
	fpsR := int(c.ws.Row)
	c.writeCursor(fpsR-2, 0) // this is basically the line of the last cell column (because it's exclusive values)
	_, _ = c.drawBuf.WriteString("--FPS")
	c.writeCursor(fpsR-1, 0)
	_, err := c.drawBuf.WriteString(fmt.Sprintf("FPS: %.2f", fps))
	if err != nil {
		logger.Error("Error printing engine stats %s", err.Error())
	}
}

// todo: this should probably be done a folder for the game. The engine should just be worried about running the game
// the canvas should know what to do with which piece
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
						pixelR: i,
						pixelC: j,
					},
					Valid: false,
				}
			} else {
				// we don't support special blocks, so we assume everything is valid for now
				cells[i][j] = Cell{
					Color: colors.Black,
					Position: PixelPoint{
						pixelR: i,
						pixelC: j,
					},
					Valid: true,
				}
			}

		}
	}

	return cells
}
