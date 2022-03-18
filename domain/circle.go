package domain

type Circle struct {
	Center Coordinate
	Radius int
}

func NewCircle(center Coordinate, radius int) *Circle {
	return &Circle{
		Center: center,
		Radius: radius,
	}
}

func (c *Circle) GetCoordinates() Coordinates {
	var coordinates Coordinates
	initialCoordinates := c.getInitialCoordinates()

	for _, point := range initialCoordinates {

		coors := Coordinates{
			// quadrant 1
			Coordinate{X: point.X + c.Center.X, Y: point.Y + c.Center.Y},
			Coordinate{X: point.Y + c.Center.X, Y: point.X + c.Center.Y},
			// quadrant 2
			Coordinate{X: point.X + c.Center.X, Y: -point.Y + c.Center.Y},
			Coordinate{X: point.Y + c.Center.X, Y: -point.X + c.Center.Y},
			// quadrant 3
			Coordinate{X: -point.X + c.Center.X, Y: -point.Y + c.Center.Y},
			Coordinate{X: -point.Y + c.Center.X, Y: -point.X + c.Center.Y},
			// quadrant 4
			Coordinate{X: -point.X + c.Center.X, Y: point.Y + c.Center.Y},
			Coordinate{X: -point.Y + c.Center.X, Y: point.X + c.Center.Y},
		}
		coordinates = append(coordinates, coors...)
	}

	return coordinates
}

// getInitialCoordinates calculates the initial coordinates
func (c *Circle) getInitialCoordinates() Coordinates {
	var coordinates Coordinates

	point := Coordinate{0, c.Radius} // first coordinate
	p := 5/4 - c.Radius              // decision parameter

	for point.X <= point.Y {
		coordinates = append(coordinates, Coordinate{
			X: point.X,
			Y: point.Y,
		})

		if p < 0 {
			point.X++
			p = p + 2*point.X + 1
		} else {
			point.X++
			point.Y--
			p = p + 2*point.X + 1 - 2*point.Y
		}
	}

	return coordinates
}
