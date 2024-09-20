package game

import (
	"raidline/space-invaders/game/cell"
	"raidline/space-invaders/pkg/assert"
	"raidline/space-invaders/pkg/logger"
)

const (
	shipSideSize         int = 5
	enemyLength              = 3
	topEnemyLine             = 2
	enemyMargin              = 10
	enemyWidth               = 7
	enemyDescendingTicks     = 10
)

type Game struct {
	Board       [][]BoardPoint
	TickChan    <-chan bool
	enemies     []*cell.Enemy
	ship        *cell.Ship
	enemyTicker int
}

type BoardPoint struct {
	Valid bool //just to indicate if valid position
	C     cell.Cell
}

func Make(rows, cols uint16, tickChan <-chan bool) *Game {
	r := int(rows)
	c := int(cols)
	b := constructBoard(r, c)

	enemies := createEnemies(c)
	ship := cell.MakeShip(fillShipPositions(shipSideSize, r-1, c))

	placeCellsInBoard(b, enemies, ship)

	g := &Game{
		Board:       b,
		enemies:     enemies,
		ship:        ship,
		TickChan:    tickChan,
		enemyTicker: 0,
	}

	go g.enemyMoverTicker()

	return g
}

func (g *Game) enemyMoverTicker() {
	for tick := range g.TickChan {

		assert.Assert(tick, "We should never receive false for the tick channel")

		g.enemyTicker++
		if g.enemyTicker == enemyDescendingTicks {
			g.enemyTicker = 0
			for _, enemy := range g.enemies {
				cannonShipPos := g.ship.Positions[0]

				//todo: signal game end
				lastR := enemy.Positions[len(enemy.Positions)-1].PixelR + 1

				if lastR == g.ship.Positions[0].PixelR {
					logger.Error("GAME ENDED %s", "shame")
				}

				newPositions := make([]*cell.PixelPoint, len(enemy.Positions))
				for i, position := range enemy.Positions {
					prevR := position.PixelR
					prevC := position.PixelC

					g.Board[prevR][prevC] = createEmptyPoint()

					newPositions[i] = &cell.PixelPoint{
						PixelR: prevR + 1,
						PixelC: prevC,
					}
				}

				enemy.Positions = nil
				enemy.Positions = newPositions

				for _, position := range enemy.Positions {
					assert.Assert(position.PixelR < cannonShipPos.PixelR, "You cannot move past the cannon of the ship Line %d", cannonShipPos.PixelR)
					g.Board[position.PixelR][position.PixelC] = BoardPoint{
						Valid: true,
						C:     enemy,
					}
				}

			}
		}
	}
}

func createEnemies(cols int) []*cell.Enemy {
	requiredSpace := (enemyMargin * 2) + (enemyWidth * 2)

	assert.Assert(cols > requiredSpace, "Not enough space to play the game, you should have at least %d columns to spawn 2 enemies", requiredSpace)

	enemies := make([]*cell.Enemy, 0, 2)

	for i := 0; true; i++ {

		startPoint := enemyMargin + ((enemyWidth + enemyMargin) * i)
		endPoint := startPoint + enemyWidth

		if endPoint >= cols-enemyMargin {
			break
		}

		positions := fillEnemyPositions(enemyLength, enemyWidth, topEnemyLine, startPoint)
		enemies = append(enemies, cell.MakeEnemy(positions))
	}

	return enemies
}

func placeCellsInBoard(b [][]BoardPoint, enemies []*cell.Enemy, ship *cell.Ship) {

	for _, enemy := range enemies {
		assert.NonNil(enemy)
		for _, position := range enemy.Positions {
			b[position.PixelR][position.PixelC] = BoardPoint{
				Valid: true,
				C:     enemy,
			}
		}
	}

	for _, position := range ship.Positions {
		assert.NonNil(position)
		b[position.PixelR][position.PixelC] = BoardPoint{
			Valid: true,
			C:     ship,
		}
	}
}

func fillEnemyPositions(size, enemyWidth, top, startPoint int) []*cell.PixelPoint {
	pixelPoints := make([]*cell.PixelPoint, size*(enemyWidth))
	pointsIdx := 0

	for i := 0; i < size; i++ {
		for j := 0; j < enemyWidth; j++ {
			pixelPoints[pointsIdx] = &cell.PixelPoint{
				PixelR: top + i,
				PixelC: j + startPoint,
			}
			pointsIdx++
		}
	}

	return pixelPoints
}

func fillShipPositions(sideSize, rows, cols int) []*cell.PixelPoint {
	positions := make([]*cell.PixelPoint, (sideSize*2)+(sideSize*sideSize)+1)
	posIdx := 0
	midPoint := (cols / 2) - 1
	lowerPoint := rows - 2 // one line space from the end

	//cannon
	cannonR := lowerPoint - sideSize
	positions[posIdx] = &cell.PixelPoint{
		PixelR: cannonR,
		PixelC: midPoint,
	}

	posIdx++

	//rows
	lowerSize := 0
	for i := 1; i <= sideSize; i++ {
		if lowerSize == 0 {
			lowerSize = 3
		} else {
			lowerSize = lowerSize + 2
		}
		newRowR := positions[posIdx-1].PixelR + 1
		newRowC := (midPoint - i) - 1
		for j := 1; j <= lowerSize; j++ {
			positions[posIdx] = &cell.PixelPoint{
				PixelR: newRowR,
				PixelC: newRowC + j,
			}

			posIdx++
		}
	}

	return positions
}

func constructBoard(rows, cols int) [][]BoardPoint {
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
				outerBoard[i][j] = createEmptyPoint()
			}
		}
	}

	return outerBoard
}

func createEmptyPoint() BoardPoint {
	return BoardPoint{
		Valid: true,
		C:     nil,
	}
}

// could be a "static" method (could help in collision detection)
func isBorder(rows int, cols int, j int, i int) bool {
	return j == 0 || i == 0 || j == cols-1 || i == rows-1
}
