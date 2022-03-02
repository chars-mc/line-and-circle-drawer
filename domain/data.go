package domain

type Coordinate struct {
	x int
	y int
}

type Coordinates []Coordinate

type Line interface {
	GetCoordinates() Coordinates
}
