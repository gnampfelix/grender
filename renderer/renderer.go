package renderer

import (
	"github.com/gnampfelix/grender/geometry"
)

type Renderer interface {
	Render(input Input) Output
}

type gnampfelixRenderer struct {
	output Output
}

func New(output Output) Renderer {
	return gnampfelixRenderer{output:output}
}


func (g gnampfelixRenderer) Render(input Input) Output {
	width := 480
	height := 270
	output := g.output
	depth := NewMapBuffer()

	camera := NewCamera(geometry.NewVector3(0, -20, 0), 2)
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
			for input.HasNextObject() {
				currentObject := input.NextObject()
				for currentObject.HasNextTriangle() {
					currentTri := currentObject.NextTriangle()
					if currentTri.IsHit(&ray) {
						hitPoint, _ := ray.HitPoint()
						distance := geometry.Subtract(hitPoint, camera.Origin()).Length()
						if depth.SetDepthIfCloser(distance, x, y) {
							output.SetPixel(currentTri.Color(), x, y)
						}
					}
				}
				currentObject.Reset()
			}
			input.Reset()
		}
	}
	return output
}
