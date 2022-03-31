package domain

import (
	"errors"
	"fmt"
	"strings"
)

type Polygon struct {
	coordinates Coordinates
}

func NewPolygon(coordinates Coordinates) *Polygon {
	return &Polygon{
		coordinates: coordinates,
	}
}

func (p *Polygon) GetCoordinates() Coordinates {
	return p.coordinates
}

func (p *Polygon) String() string {
	return fmt.Sprint(p.coordinates)
}

func (p *Polygon) Set(coord string) error {
	values := strings.Split(coord, ";")

	if len(values) < 3 {
		return errors.New("not enough coordinates, at least 3 are nedded")
	}

	for _, v := range values {
		var coordinate Coordinate
		err := coordinate.Set(v)
		if err != nil {
			return err
		}
		p.coordinates = append(p.coordinates, coordinate)
	}

	return nil
}
