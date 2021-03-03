package renderer

import (
	"math"

	"github.com/gnampfelix/grender/geometry"
)

type Object interface {
	HasNextTriangle() bool
	NextTriangle() geometry.Triangle
	Transform(matrix geometry.Matrix44)
	Rotate(axis geometry.Axis, angle float64)
	Reset()
}

type cubeObject struct {
	triangles            []geometry.Triangle
	current              int
	transformedTriangles []geometry.Triangle
	transformationMatrix geometry.Matrix44
}

func NewCube() Object {
	tris := make([]geometry.Triangle, 0)
	transformedTriangles := make([]geometry.Triangle, 0)
	for i := 0; i < (len(triangles) / 4); i++ {
		currentTri := geometry.NewTriangleWithName(triangles[i*4], triangles[i*4+1], triangles[i*4+2], triangles[i*4+3], names[i])
		tris = append(tris, currentTri)
		transformedTriangles = append(transformedTriangles, currentTri)
	}

	return &cubeObject{
		triangles:            tris,
		current:              0,
		transformedTriangles: transformedTriangles,
		transformationMatrix: geometry.NewMatrix44(
			geometry.NewVector4(1, 0, 0, 0),
			geometry.NewVector4(0, 1, 0, 0),
			geometry.NewVector4(0, 0, 1, 0),
			geometry.NewVector4(0, 0, 0, 1),
		),
	}
}

func (c cubeObject) HasNextTriangle() bool {
	return c.current < len(c.triangles)
}

func (c *cubeObject) NextTriangle() geometry.Triangle {
	if c.HasNextTriangle() {
		tri := c.transformedTriangles[c.current]
		c.current++
		return tri
	}
	return geometry.Triangle{}
}

func (c *cubeObject) Reset() {
	c.current = 0
}

func (c *cubeObject) Transform(matrix geometry.Matrix44) {
	c.transformationMatrix = matrix
	for index := range c.triangles {
		c.transformedTriangles[index] = c.triangles[index].Transform(matrix)
	}
}

func (c *cubeObject) Rotate(axis geometry.Axis, angle float64) {
	convertedAngle := geometry.DegToRad(angle)
	cosAngle := math.Cos(convertedAngle)
	sinAngle := math.Sin(convertedAngle)

	var transformationMatrix geometry.Matrix44

	switch axis {
	case geometry.X:
		transformationMatrix = geometry.NewMatrix44(
			geometry.NewVector4(1, 0, 0, 0),
			geometry.NewVector4(0, cosAngle, 0-sinAngle, 0),
			geometry.NewVector4(0, sinAngle, cosAngle, 0),
			geometry.NewVector4(0, 0, 0, 1),
		)
	case geometry.Y:
		transformationMatrix = geometry.NewMatrix44(
			geometry.NewVector4(cosAngle, 0, sinAngle, 0),
			geometry.NewVector4(0, 1, 0, 0),
			geometry.NewVector4(0-sinAngle, 0, cosAngle, 0),
			geometry.NewVector4(0, 0, 0, 1),
		)
	case geometry.Z:
		transformationMatrix = geometry.NewMatrix44(
			geometry.NewVector4(cosAngle, 0-sinAngle, 0, 0),
			geometry.NewVector4(sinAngle, cosAngle, 0, 0),
			geometry.NewVector4(0, 0, 1, 0),
			geometry.NewVector4(0, 0, 0, 1),
		)
	}
	c.Transform(transformationMatrix.Chain(c.transformationMatrix))
}

var names = []string{
	"front", "front", "unten", "unten", "rechts", "rechts", "links", "links", "oben", "oben", "front", "front", "unten", "unten", "rechts", "rechts", "links", "links", "oben", "oben",
}

var triangles = []geometry.Vector3{
	geometry.NewVector3(-5, -10, -5), geometry.NewVector3(-5, 0, -5), geometry.NewVector3(5, -10, -5), geometry.NewVector3(191, 97, 106), //red, front
	//geometry.NewVector3(5, -10, -5), geometry.NewVector3(-5, 0, -5), geometry.NewVector3(5, 0, -5), geometry.NewVector3(191, 97, 106), //red

	geometry.NewVector3(5, -10, -5), geometry.NewVector3(5, -10, -15), geometry.NewVector3(5, 0, -5), geometry.NewVector3(143, 188, 187), //blueish, right side
	geometry.NewVector3(5, 0, -5), geometry.NewVector3(5, 0, -15), geometry.NewVector3(5, -10, -15), geometry.NewVector3(143, 188, 187), //blueish, right side

	geometry.NewVector3(-5, -10, -5), geometry.NewVector3(-5, -10, -15), geometry.NewVector3(5, -10, -5), geometry.NewVector3(236, 239, 244), //white, bottom
	geometry.NewVector3(-5, -10, -15), geometry.NewVector3(5, -10, -15), geometry.NewVector3(5, -10, -5), geometry.NewVector3(236, 239, 244), //white, bottom

	geometry.NewVector3(-5, -10, -15), geometry.NewVector3(-5, 0, -15), geometry.NewVector3(5, -10, -15), geometry.NewVector3(191, 97, 106), //red, back
	geometry.NewVector3(5, -10, -15), geometry.NewVector3(-5, 0, -15), geometry.NewVector3(5, 0, -15), geometry.NewVector3(191, 97, 106), //red

	geometry.NewVector3(-5, -10, -5), geometry.NewVector3(-5, -10, -15), geometry.NewVector3(-5, 0, -5), geometry.NewVector3(143, 188, 187), //blueish, left side
	geometry.NewVector3(-5, 0, -5), geometry.NewVector3(-5, 0, -15), geometry.NewVector3(-5, -10, -15), geometry.NewVector3(143, 188, 187), //blueish, left side

	geometry.NewVector3(-5, 0, -5), geometry.NewVector3(-5, 0, -15), geometry.NewVector3(5, 0, -5), geometry.NewVector3(236, 239, 244), //white, top
	geometry.NewVector3(-5, 0, -15), geometry.NewVector3(5, 0, -15), geometry.NewVector3(5, 0, -5), geometry.NewVector3(236, 239, 244), //white, top

	//geometry.NewVector3(-4, 0, -5), geometry.NewVector3(6, 0, -5), geometry.NewVector3(-4, 0, -10), geometry.NewVector3(143, 188, 187), //blueish
	//geometry.NewVector3(-5, -5, -5), geometry.NewVector3(-5, 5, -5), geometry.NewVector3(5, -5, -5), geometry.NewVector3(143, 188, 187), //blueish

	//geometry.NewVector3(-5, 5, -5), geometry.NewVector3(5, 5, -5), geometry.NewVector3(-5, 5, 5), geometry.NewVector3(143, 188, 187), //blueish
	//geometry.NewVector3(5, 5, -5), geometry.NewVector3(-5, 5, 5), geometry.NewVector3(5, 5, 5), geometry.NewVector3(143, 188, 187), //blueish

}
