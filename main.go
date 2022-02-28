package main

import (
	"runtime"

	"github.com/chars-mc/opengl-exercises/ui"
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

	for !ui.Window.ShouldClose() {
		ui.Draw()
	}
}
