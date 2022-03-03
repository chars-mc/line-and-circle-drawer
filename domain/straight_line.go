package domain

import "math"

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

	m := float64(line.b.Y-line.a.Y) / float64(line.b.X-line.a.X) // get the slope of the line
	b := float64(line.a.Y) - m*float64(line.a.X)                 // find b
	b = math.Round(b*10) / 10                                    // round b to one decimals
	k := line.b.X - line.a.X                                     // get number of iterations

	for i := 0; i < k; i++ {
		y := m*float64(line.a.X) + b // yk = mx + b

		coordinates = append(coordinates, Coordinate{
			X: line.a.X,
			Y: int(math.Round(y)),
		})
		line.a.X++
	}

	return coordinates
}
