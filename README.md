# Line and circle drawer

Implementation of algorithms for drawing lines and circles with OpenGL.

## Requirements
* [gcc](https://www.msys2.org/)
* [Go](https://go.dev/)

## Libraries
* [gl](https://github.com/go-gl/gl)
* [glfw](https://github.com/go-gl/glfw)

## Instructions
1. Install  [Go](https://go.dev/)
2. Install [MSYS2](https://www.msys2.org/)
3. Execute `github.com/go-gl/gl/v4.4-compatibility/gl`
4. Build the project: `make build` or `go build main.go`

## Resources
[GLFW Documentation](https://www.glfw.org/documentation.html)

## Usage
### Line
```
λ opengl-exercises.exe line -h
Usage of line:
  -from value
        the start coordinate
  -to value
        the end coordinate  
```

### Circle
```
λ opengl-exercises.exe circle -h 
Usage of circle:
  -center value
        the center of the circumference
  -radius int
        the radius of the circumference (default 10)
```
