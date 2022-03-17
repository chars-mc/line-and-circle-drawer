package domain

type Circle struct {
	Center Coordinate
	Radio  int
}

func NewCircle(center Coordinate, radio int) *Circle {
	return &Circle{
		Center: center,
		Radio:  radio,
	}
}

func (c *Circle) GetCoordinates() Coordinates {
	var coordinates Coordinates

	// TODO: implement algorithm

	return coordinates
}
