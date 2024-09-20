package cell

type Cell interface {
	Print(i, j int) string
	IsFoundInPositions(i, j int) bool //todo: we are doing for loops.. maybe we could have a map of some kind to make the search faster?
}

type PixelPoint struct {
	PixelR int
	PixelC int
}

type Ship struct {
	Positions []*PixelPoint // the first position is of the cannon
}

func MakeShip(positions []*PixelPoint) *Ship {
	return &Ship{
		Positions: positions,
	}
}

// Print gives the correct char for the position, in order to be drawn
func (s *Ship) Print(i, j int) string {
	return "@"
}

func (s *Ship) IsFoundInPositions(i int, j int) bool {
	for _, position := range s.Positions {
		if position.PixelR == i && position.PixelC == j {
			return true
		}
	}

	return false
}

// We could have multiple types of enemies in the future

type Enemy struct {
	Positions []*PixelPoint
}

func MakeEnemy(positions []*PixelPoint) *Enemy {
	return &Enemy{
		Positions: positions,
	}
}

// Print gives the correct char for the position, in order to be drawn
func (s *Enemy) Print(i, j int) string {

	lastPoint := s.Positions[len(s.Positions)-1]
	firstPoint := s.Positions[0]

	columnBegin := firstPoint.PixelC
	columnEnd := lastPoint.PixelC
	width := columnEnd - columnBegin

	rowBegin := firstPoint.PixelR
	rowEnd := lastPoint.PixelR
	length := rowEnd - rowBegin

	if (i - rowBegin) == 1 {
		if j-columnBegin == (width/2)-1 {
			return "O"
		}

		if j-columnBegin == (width/2)+1 {
			return "O"
		}
	}

	if (i-rowBegin) == length && j-columnBegin == (width/2) {
		return "□"
	}

	return "."
}

func (s *Enemy) IsFoundInPositions(i int, j int) bool {
	for _, position := range s.Positions {
		if position.PixelR == i && position.PixelC == j {
			return true
		}
	}

	return false
}
