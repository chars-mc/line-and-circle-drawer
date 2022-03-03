package main

import (
	"log"
	"runtime"

	"github.com/chars-mc/opengl-exercises/ui"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	points, err := ui.GetPoints()
	if err != nil {
		log.Fatal(err)
	}

	ui, err := ui.NewUI(640, 480, "Testing")
	if err != nil {
		panic(err)
	}
	defer ui.Terminate()

	vao := ui.GenerateVao(points)

	for !ui.Window.ShouldClose() {
		ui.Draw(vao, len(points))
	}
}
