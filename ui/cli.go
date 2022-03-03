package ui

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/chars-mc/opengl-exercises/domain"
	"github.com/chars-mc/opengl-exercises/utils"
)

var (
	from = flag.String("from", "2,2", "Start coordinate")
	to   = flag.String("to", "10,10", "End coordinate")
)

func init() {
	flag.Parse()
}

// GetPoints proccess the input from command arguments
func GetPoints() ([]float32, error) {
	first := strings.Split(*from, ",")
	last := strings.Split(*to, ",")

	if len(os.Args) == 1 {
		log.Println("coordinates not provided, using default values")
	}

	if len(first) < 2 || len(last) < 2 {
		return nil, errors.New("Wrong coordinates")
	}

	x1, err := strconv.Atoi(first[0])
	if err != nil {
		return nil, fmt.Errorf("Error on x1(%v): %v", first[0], err)
	}
	y1, err := strconv.Atoi(first[1])
	if err != nil {
		return nil, fmt.Errorf("Error on y1(%v): %v", first[1], err)
	}
	x2, err := strconv.Atoi(last[0])
	if err != nil {
		return nil, fmt.Errorf("Error on x2(%v): %v", last[0], err)
	}
	y2, err := strconv.Atoi(last[1])
	if err != nil {
		return nil, fmt.Errorf("Error on x2(%v): %v", last[1], err)
	}

	a := domain.Coordinate{X: x1, Y: y1}
	b := domain.Coordinate{X: x2, Y: y2}

	coordinates := domain.NewStraightLine(a, b).GetCoordinates()

	return utils.GetPointsFromCoordinates(coordinates), nil
}
