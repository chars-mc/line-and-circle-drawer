package ui

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type UI struct {
	Window  *glfw.Window
	Program uint32
}

// NewUI initializes glfw and gl
func NewUI(width, height int, title string) (*UI, error) {
	if err := glfw.Init(); err != nil {
		return nil, err
	}

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return nil, err
	}

	window.MakeContextCurrent()

	if err = gl.Init(); err != nil {
		return nil, err
	}

	program := gl.CreateProgram()
	gl.LinkProgram(program)

	return &UI{
		Window:  window,
		Program: program,
	}, nil
}
