package renderer

import (
	"math"

	"github.com/gnampfelix/grender/geometry"
)

type rasterizationRenderer struct {
}

func NewRasterizationRenderer() Renderer {
	return rasterizationRenderer{}
}

func (g rasterizationRenderer) Render(input Input, output Output) {
	width := output.Width()
	height := output.Height()
	depth := NewMapBuffer()
	nearClippingPlane := 1.0

	//world := geometry.NewVector3(0, 0, 0)
	//camera := NewCamera(geometry.NewVector3(0, -20, 0), 2)
	//screen := camera.Screen()

	for input.HasNextObject() {
		currentObject := input.NextObject()
		for currentObject.HasNextTriangle() {
			currentTri := currentObject.NextTriangle()
			currentTri.DividePerspective(nearClippingPlane)
			currentTri.MapToNdcSpace(2, 3, -3, -4)
			currentTri.MapToRasterSpacer(height, width)

			minX := int(math.Min(math.Max(0, math.Min(currentTri.A().X(), math.Min(currentTri.B().X(), currentTri.C().X()))), float64(width-1)))
			maxX := int(math.Max(math.Min(float64(width-1), math.Max(currentTri.A().X(), math.Max(currentTri.B().X(), currentTri.C().X()))), 0))

			minY := int(math.Min(math.Max(0, math.Min(currentTri.A().Y(), math.Min(currentTri.B().Y(), currentTri.C().Y()))), float64(height-1)))
			maxY := int(math.Max(math.Min(float64(height-1), math.Max(currentTri.A().Y(), math.Max(currentTri.B().Y(), currentTri.C().Y()))), 0))
			for y := minY; y <= maxY; y++ {
				for x := minX; x <= maxX; x++ {
					a := currentTri.A().ExtractVector2()
					b := currentTri.B().ExtractVector2()
					c := currentTri.C().ExtractVector2()
					point := geometry.NewVector2(float64(x), float64(y))

					area := geometry.CalculateEdgeFunction2(a, b, c) // Area of the triangle * 2
					w0 := geometry.CalculateEdgeFunction2(b, a, point)
					w1 := geometry.CalculateEdgeFunction2(c, b, point)
					w2 := geometry.CalculateEdgeFunction2(a, c, point)

					if x > 340 && y > 100 {
						math.Abs(3)
					}
					if (w0 >= 0 && w1 >= 0 && w2 >= 0) || (w0 <= 0 && w1 <= 0 && w2 <= 0) {
						w0 /= area
						w1 /= area
						w2 /= area
						// Interpolate the depth while accounting for perpsective divide that was done earlier
						currentDepth := 1 / (currentTri.A().Z()*w0 + currentTri.B().Z()*w1 + currentTri.C().Z()*w2) // Interpolate the
						if depth.SetDepthIfCloser(currentDepth, x, y) {                                             // need barycentric coordinates for that :(
							output.SetPixel(currentTri.Color(), x, y)
						}
					}
				}
			}
		}
		currentObject.Reset()
	}
	input.Reset()
}
