package domain

type Coordinate struct {
	X int
	Y int
}

type Coordinates []Coordinate

type Line interface {
	GetCoordinates() Coordinates
}
