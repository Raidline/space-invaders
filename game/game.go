package game

import "raidline/space-invaders/game/cell"

const (
	shipSideSize  int = 2
	shipLowerSize int = 1
)

type Game struct {
	Board [][]BoardPoint
	// we should probably also have a relation to position Cell -> so we can call it to Print
	ship *cell.Ship
}

type BoardPoint struct {
	Valid bool //just to indicate if valid position
	C     *cell.Cell
}

func Make(rows, cols uint16) *Game {
	ship := cell.MakeShip(shipSideSize, shipLowerSize, int(rows), int(cols))

	//todo: stats takes 2 lines (we should have this calculated)
	b := constructBoard(int(rows-2), int(cols), ship)

	return &Game{
		Board: b,
		ship:  ship,
	}
}

func constructBoard(rows, cols int, ship *cell.Ship) [][]BoardPoint {
	outerBoard := make([][]BoardPoint, rows)
	for i := 0; i < rows; i++ {
		outerBoard[i] = make([]BoardPoint, cols)
		for j := 0; j < cols; j++ {
			if isBorder(rows, cols, j, i) {
				outerBoard[i][j] = BoardPoint{
					Valid: false,
					C:     nil,
				}
			} else {
				outerBoard[i][j] = BoardPoint{
					Valid: true,
					C:     nil,
				}
			}
		}
	}

	return outerBoard
}

// could be a "static" method (could help in collision detection)
func isBorder(rows int, cols int, j int, i int) bool {
	return j == 0 || i == 0 || j == cols-1 || i == rows-1
}
