package ui

import (
	"github.com/go-gl/gl/v4.4-compatibility/gl"
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

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCompatProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

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

func (ui *UI) Draw(vao uint32, length int) {
	gl.UseProgram(ui.Program)
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.PointSize(2)
	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.POINTS, 0, int32(length/3))

	glfw.PollEvents()
	ui.Window.SwapBuffers()
}

// GenerateVao initializes and returns a vertex array object from points
func (ui *UI) GenerateVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func (ui *UI) Terminate() {
	glfw.Terminate()
}
