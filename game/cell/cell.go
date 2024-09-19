package cell

type Ship struct {
	positions []*pixelPoint // the first position is of the cannon
}

func MakeShip(sideSize, lowerSize, rows, cols int) *Ship {
	return &Ship{
		positions: fillShipPositions(sideSize, lowerSize, rows, cols),
	}
}

//should be in canvas??

// Print gives the correct char for the position, in order to be drawn
func (s *Ship) Print(i, j int) {

}

type Cell interface {
	Print(i, j int)
}

type pixelPoint struct {
	pixelR int
	pixelC int
}

func fillShipPositions(sideSize, lowerSize, rows, cols int) []*pixelPoint {
	positions := make([]*pixelPoint, (sideSize*2)+(lowerSize*4)+1)
	posIdx := 0
	midPoint := cols / 2
	lowerPoint := rows - 5

	//cannon
	cannonR := lowerPoint - sideSize
	positions[posIdx] = &pixelPoint{
		pixelR: cannonR,
		pixelC: midPoint,
	}

	posIdx++

	//left side
	for delta := 1; delta <= sideSize; delta++ {
		positions[posIdx] = &pixelPoint{
			pixelR: cannonR + delta,
			pixelC: midPoint - delta,
		}
		posIdx++
	}

	// right side
	for delta := 1; delta <= sideSize; delta++ {
		positions[posIdx] = &pixelPoint{
			pixelR: cannonR + delta,
			pixelC: midPoint + delta,
		}
		posIdx++
	}

	lastPosIdx := posIdx - 1
	// lower left
	for delta := 1; delta <= lowerSize; delta++ {
		positions[posIdx] = &pixelPoint{
			pixelR: positions[lastPosIdx].pixelR,
			pixelC: positions[lastPosIdx].pixelC + delta,
		}
		posIdx++
	}

	lastPosIdx = posIdx - 1
	//lower left mid
	for delta := 1; delta <= lowerSize; delta++ {
		positions[posIdx] = &pixelPoint{
			pixelR: positions[lastPosIdx].pixelR - delta,
			pixelC: positions[lastPosIdx].pixelC + delta,
		}
		posIdx++
	}
	//lower right mid
	lastPosIdx = posIdx - 1
	for delta := 1; delta <= lowerSize; delta++ {
		positions[posIdx] = &pixelPoint{
			pixelR: positions[lastPosIdx].pixelR + delta,
			pixelC: positions[lastPosIdx].pixelC + delta,
		}
		posIdx++
	}

	//lower right
	lastPosIdx = posIdx - 1
	for delta := 1; delta <= lowerSize; delta++ {
		positions[posIdx] = &pixelPoint{
			pixelR: positions[lastPosIdx].pixelR,
			pixelC: positions[lastPosIdx].pixelC + delta,
		}
		posIdx++
	}

	return positions
}
