package main

import (
	"runtime"

	"github.com/chars-mc/opengl-exercises/ui"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	ui, err := ui.NewUI(640, 480, "Testing")
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	for !ui.Window.ShouldClose() {
		ui.Window.SwapBuffers()
		glfw.PollEvents()
	}
}
