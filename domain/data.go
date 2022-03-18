package domain

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("%v,%v", c.X, c.Y)
}

func (c *Coordinate) Set(coord string) error {
	values := strings.Split(coord, ",")

	if len(values) < 2 {
		return errors.New("wrong coordinate params")
	}

	x, err := strconv.Atoi(values[0])
	if err != nil {
		return errors.New("wrong value for x")
	}

	y, err := strconv.Atoi(values[1])
	if err != nil {
		return errors.New("wrong value for y")
	}

	c.X = x
	c.Y = y
	return nil
}

type Coordinates []Coordinate

type Line interface {
	GetCoordinates() Coordinates
}
