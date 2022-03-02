package domain

// StraightLine represents a line from coordinates a to b
type StraightLine struct {
	a Coordinate
	b Coordinate
}

func NewStraightLine(a, b Coordinate) *StraightLine {
	return &StraightLine{
		a: a,
		b: b,
	}
}

func (line *StraightLine) GetCoordinates() Coordinates {
	var coordinates Coordinates

	// TODO: implement algorithm

	return coordinates
}
