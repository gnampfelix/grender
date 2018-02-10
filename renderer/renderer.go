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
	width := 1000
	height := 1000

	output := NewSimpleOutput(height, width)
	depth := NewMapBuffer()

	camera := geometry.NewVector3(6, 0, 10)
	leftBottom := geometry.NewVector3(4, 2, 12)
	leftTop := geometry.NewVector3(4, 2, 8)
	rightBottom := geometry.NewVector3(8, 2, 12)
	//rightTop := geometry.NewVector3(8, 2, 8)

	//bottom := geometry.NewRay(leftBottom, geometry.Subtract(rightBottom, leftBottom))
	//left := geometry.NewRay(leftBottom, geometry.Subtract(leftTop, leftBottom))

	screen := geometry.NewPlane(leftBottom, geometry.Subtract(rightBottom, leftBottom), geometry.Subtract(leftTop, leftBottom))

	lineLength := geometry.Subtract(rightBottom, leftBottom).Length()
	colLength := geometry.Subtract(leftTop, leftBottom).Length()

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

			ray := geometry.NewRay(camera, geometry.Subtract(screenPoint, camera))
			for input.HasNextTriangle() {
				currentTri := input.NextTriangle()
				if currentTri.IsHit(&ray) {
					hitPoint, _ := ray.HitPoint()
					distance := geometry.Subtract(hitPoint, camera).Length()
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
