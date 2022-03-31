package ui

import (
	"github.com/chars-mc/opengl-exercises/domain"
	"github.com/chars-mc/opengl-exercises/utils"
	"github.com/go-gl/gl/v4.4-compatibility/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type UI struct {
	Window      *glfw.Window
	Program     uint32
	coordinates domain.Coordinates
	points      []float32
	vao         uint32
}

// NewUI initializes glfw and gl
func NewUI(width, height int, title string, coordinates domain.Coordinates) (*UI, error) {
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
	gl.Viewport(0, 0, int32(width), int32(height))
	ui := &UI{
		Window:      window,
		Program:     program,
		coordinates: coordinates,
	}
	glfw.GetCurrentContext().SetKeyCallback(ui.handleKeyboard)

	return ui, nil
}

func (ui *UI) Draw() {
	gl.UseProgram(ui.Program)
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.PointSize(2)

	glfw.PollEvents()
}

// GenerateVao initializes and returns a vertex array object from points
func (ui *UI) generateVao() uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(ui.points), gl.Ptr(ui.points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func (ui *UI) handleKeyboard(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	// TODO: move transformacion logic to domain layer
	// translation
	if key == glfw.KeyW && action == glfw.Press {
		for i := range ui.coordinates {
			ui.coordinates[i].Y++
		}
	} else if key == glfw.KeyS && action == glfw.Press {
		for i := range ui.coordinates {
			ui.coordinates[i].Y--
		}
	}
	if key == glfw.KeyA && action == glfw.Press {
		for i := range ui.coordinates {
			ui.coordinates[i].X--
		}
	} else if key == glfw.KeyD && action == glfw.Press {
		for i := range ui.coordinates {
			ui.coordinates[i].X++
		}
	}

	// rotation
	if key == glfw.KeyR && action == glfw.Press {
		for i := range ui.coordinates {
			ui.coordinates[i].X, ui.coordinates[i].Y = ui.coordinates[i].Y, -ui.coordinates[i].X
		}
	}

	// scalation
	// FIXME: fix coordinate changes
	if key == glfw.KeyUp && action == glfw.Press {
		for i := range ui.coordinates {
			if ui.coordinates[i].X > 0 {
				ui.coordinates[i].X++
			} else {
				ui.coordinates[i].X--
			}
			if ui.coordinates[i].Y > 0 {
				ui.coordinates[i].Y++
			} else {
				ui.coordinates[i].Y--
			}
		}
	} else if key == glfw.KeyDown && action == glfw.Press {
		for i := range ui.coordinates {
			if ui.coordinates[i].X > 0 {
				ui.coordinates[i].X--
			} else {
				ui.coordinates[i].X++
			}
			if ui.coordinates[i].Y > 0 {
				ui.coordinates[i].Y--
			} else {
				ui.coordinates[i].Y++
			}
		}
	}

	ui.points = utils.GetPointsFromCoordinates(ui.coordinates)
	ui.vao = ui.generateVao()

	gl.BindVertexArray(ui.vao)
	gl.DrawArrays(gl.LINE_LOOP, 0, int32(len(ui.points)/3))
	ui.Window.SwapBuffers()
}

func (ui *UI) Terminate() {
	glfw.Terminate()
}
