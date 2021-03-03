package main

import (
	"github.com/gnampfelix/grender/geometry"
	"github.com/gnampfelix/grender/renderer"
	"github.com/gonutz/prototype/draw"
)

const height int = 270
const width int = 480

var object renderer.Object
var input renderer.Input
var rend renderer.Renderer
var output renderer.Output

type liveOutput struct {
	renderer.Output
	window draw.Window
}

func (l liveOutput) SetPixel(c geometry.Vector3, x, y int) {
	l.window.DrawPoint(x, y, draw.RGB(float32(c.X()/255), float32(c.Y()/255), float32(c.Z()/255)))
}

func main() {
	object = renderer.NewCube()
	input = renderer.NewInput()
	input.Add(object)
	rend = renderer.NewRasterizationRenderer()
	output = renderer.NewSimpleOutput(height, width)
	draw.RunWindow("grender", width, height, update)
}

func update(window draw.Window) {
	object.Rotate(geometry.Z, 1.2)
	rend.Render(input, liveOutput{output, window})
}
