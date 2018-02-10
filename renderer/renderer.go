package renderer

import (
	"github.com/gnampfelix/grender/geometry"
)

type Renderer interface {
	Render(input Input) Output
}

type gnampfelixRenderer bool

func New() Renderer {
	return gnampfelixRenderer(true)
}

func (g gnampfelixRenderer) Render(input Input) Output {
	width := 1600
	height := 900
	output := NewSimpleOutput(height, width)
	depth := NewMapBuffer()

	camera := NewCamera(geometry.NewVector3(20, -10, 10), 1.15)
	screen := camera.Screen()
	lineLength := camera.LineLength()
	colLength := camera.ColLength()

	pixelLineStep := lineLength / float64(width)
	pixelColStep := colLength / float64(height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			screenPoint := screen.P()
			screenU := screen.U()
			screenV := screen.V()

			screenU.MultiplyWithScalar(float64(x) * pixelLineStep)
			screenV.MultiplyWithScalar(float64(y) * pixelColStep)

			screenPoint.Add(screenU)
			screenPoint.Add(screenV)

			ray := geometry.NewRay(camera.Origin(), geometry.Subtract(screenPoint, camera.Origin()))
			for input.HasNextTriangle() {
				currentTri := input.NextTriangle()
				if currentTri.IsHit(&ray) {
					hitPoint, _ := ray.HitPoint()
					distance := geometry.Subtract(hitPoint, camera.Origin()).Length()
					if depth.SetDepthIfCloser(distance, x, y) {
						output.SetPixel(currentTri.Color(), x, y)
					}
				}
			}
			input.Reset()
		}
	}
	return output
}
