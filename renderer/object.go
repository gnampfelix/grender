package renderer

import (
    "github.com/gnampfelix/grender/geometry"
)

type Object interface {
    HasNextTriangle() bool
    NextTriangle() geometry.Triangle
    Transform(matrix geometry.Matrix44)
    Reset()
}

type cubeObject struct {
    triangles []geometry.Triangle
    current   int
}

func NewCube() Object {
    tris := make([]geometry.Triangle, 0)
    for i := 0; i < (len(triangles) / 4); i++ {
        currentTri := geometry.NewTriangleWithName(triangles[i*4], triangles[i*4+1], triangles[i*4+2], triangles[i*4+3], names[i])
        tris = append(tris, currentTri)
    }

    return &cubeObject{
        triangles: tris,
        current:   0,
    }
}

func (c cubeObject) HasNextTriangle() bool {
	return c.current < len(c.triangles)
}

func (c *cubeObject) NextTriangle() geometry.Triangle {
	if c.HasNextTriangle() {
		result := c.triangles[c.current]
		c.current++
		return result
	}
	return geometry.Triangle{}
}

func (c *cubeObject) Reset() {
	c.current = 0
}

func (c *cubeObject) Transform(matrix geometry.Matrix44) {

}

var names = []string{
	"front", "front", "unten", "unten", "rechts", "rechts", "links", "links", "oben", "oben",
}

var triangles = []geometry.Vector3{
	geometry.NewVector3(-5, 10, -5), geometry.NewVector3(5, 10, -5), geometry.NewVector3(5, 10, 5), geometry.NewVector3(255, 255, 0), //front
	geometry.NewVector3(-5, 10, -5), geometry.NewVector3(5, 10, 5), geometry.NewVector3(-5, 10, 5), geometry.NewVector3(255, 255, 0),
	geometry.NewVector3(-5, 10, -5), geometry.NewVector3(5, 10, -5), geometry.NewVector3(-5, 20, -5), geometry.NewVector3(255, 0, 0), //unten
	geometry.NewVector3(5, 10, -5), geometry.NewVector3(-5, 20, -5), geometry.NewVector3(5, 20, -5), geometry.NewVector3(255, 0, 0),
	geometry.NewVector3(5, 10, -5), geometry.NewVector3(5, 10, 5), geometry.NewVector3(5, 20, -5), geometry.NewVector3(0, 255, 0), //rechts
	geometry.NewVector3(5, 10, 5), geometry.NewVector3(5, 20, 5), geometry.NewVector3(5, 20, -5), geometry.NewVector3(0, 255, 0),

	geometry.NewVector3(-5, 10, -5), geometry.NewVector3(-5, 10, 5), geometry.NewVector3(-5, 20, -5), geometry.NewVector3(0, 0, 255), //links
	geometry.NewVector3(-5, 10, 5), geometry.NewVector3(-5, 20, -5), geometry.NewVector3(-5, 20, 5), geometry.NewVector3(0, 0, 255),
	geometry.NewVector3(-5, 10, 5), geometry.NewVector3(5, 10, 5), geometry.NewVector3(-5, 20, 5), geometry.NewVector3(0, 255, 255), //oben
	geometry.NewVector3(5, 10, 5), geometry.NewVector3(-5, 20, 5), geometry.NewVector3(5, 20, 5), geometry.NewVector3(0, 255, 255),
}
