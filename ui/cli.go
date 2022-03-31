package ui

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/chars-mc/opengl-exercises/domain"
)

func init() {
	flag.Parse()
}

// GetPoints proccess the input from command arguments
func GetPoints() (domain.Coordinates, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("you must to provide a command")
	}

	switch os.Args[1] {
	case "circle":
		var center domain.Coordinate

		circleCmd := flag.NewFlagSet("circle", flag.ExitOnError)
		radius := circleCmd.Int("radius", 10, "the radius of the circumference")
		circleCmd.Var(&center, "center", "the center of the circumference")

		err := circleCmd.Parse(os.Args[2:])
		if err != nil {
			return nil, err
		}

		if center == (domain.Coordinate{}) {
			center.X, center.Y = 0, 0
			log.Printf("center was not provided, using default value (%d, %d)",
				center.X, center.Y,
			)
		}

		c := domain.NewCircle(center, *radius)
		l := domain.NewStraightLine(c.Center, domain.Coordinate{c.Radius * 2, c.Radius * -2})
		coordinates := append(c.GetCoordinates(), l.GetCoordinates()...)

		return coordinates, nil
	case "line":
		var from, to domain.Coordinate

		lineCmd := flag.NewFlagSet("line", flag.ExitOnError)
		lineCmd.Var(&from, "from", "the start coordinate")
		lineCmd.Var(&to, "to", "the end coordinate")

		err := lineCmd.Parse(os.Args[2:])
		if err != nil {
			return nil, err
		}

		if from == (domain.Coordinate{}) || to == (domain.Coordinate{}) {
			from.X, from.Y = -10, -10
			to.X, to.Y = 10, 10
			log.Printf(
				"from or to coordinates are empty, using default values (%v, %d) and (%d, %d)",
				from.X, from.Y, to.X, to.Y,
			)
		}

		coordinates := domain.NewStraightLine(from, to).GetCoordinates()
		return coordinates, nil
	case "polygon":
		var polygon domain.Polygon

		polygonCmd := flag.NewFlagSet("polygon", flag.ExitOnError)
		polygonCmd.Var(&polygon, "coordinates", "the coordinates separated by a ';'")

		err := polygonCmd.Parse(os.Args[2:])
		if err != nil {
			return nil, err
		}
		if len(os.Args) < 3 {
			return nil, errors.New("you must to provide the coordinates")
		}

		return polygon.GetCoordinates(), nil
	default:
		return nil, errors.New("wrong command")
	}
}
