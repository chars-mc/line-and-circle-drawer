package main

import (
	"runtime"

	"github.com/chars-mc/opengl-exercises/domain"
	"github.com/chars-mc/opengl-exercises/ui"
	"github.com/chars-mc/opengl-exercises/utils"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	ui, err := ui.NewUI(640, 480, "Testing")
	if err != nil {
		panic(err)
	}
	defer ui.Terminate()

	// TODO: get coordinates values from stin
	coordinates := domain.NewStraightLine(domain.Coordinate{2, 2}, domain.Coordinate{100, 100}).GetCoordinates()
	points := utils.GetPointsFromCoordinates(coordinates)

	vao := ui.GenerateVao(points)

	for !ui.Window.ShouldClose() {
		ui.Draw(vao, len(points))
	}
}
