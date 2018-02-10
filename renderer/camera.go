package renderer

import (
	"github.com/gnampfelix/grender/geometry"
)

type Camera struct {
	origin                                     geometry.Vector3
	leftTop, leftBottom, rightTop, rightBottom geometry.Vector3
	screen                                     geometry.Plane
}

func NewCamera(origin geometry.Vector3, focalLength float64) Camera {
	direction := geometry.NewVector3(0, 1, 0)
	direction.MultiplyWithScalar(focalLength)
	screenCenter := geometry.Add(origin, direction)

	leftBottom := geometry.Add(screenCenter, geometry.NewVector3(-2*1.6, 0, 2*0.9))
	leftTop := geometry.Add(screenCenter, geometry.NewVector3(-2*1.6, 0, -2*0.9))
	rightBottom := geometry.Add(screenCenter, geometry.NewVector3(2*1.6, 0, 2*0.9))
	rightTop := geometry.Add(screenCenter, geometry.NewVector3(2*1.6, 0, -2*0.9))
	screen := geometry.NewPlane(leftBottom, geometry.Subtract(rightBottom, leftBottom), geometry.Subtract(leftTop, leftBottom))

	return Camera{
		origin:      origin,
		leftTop:     leftTop,
		leftBottom:  leftBottom,
		rightTop:    rightTop,
		rightBottom: rightBottom,
		screen:      screen,
	}
}

func (c Camera) LineLength() float64 {
	return geometry.Subtract(c.rightBottom, c.leftBottom).Length()
}

func (c Camera) ColLength() float64 {
	return geometry.Subtract(c.leftTop, c.leftBottom).Length()
}

func (c Camera) Screen() geometry.Plane {
	return c.screen
}

func (c Camera) Origin() geometry.Vector3 {
	return c.origin
}
