package geometry

import (
	"fmt"
)

type Triangle struct {
	a, b, c Vector3
	color   Vector3
	name    string
	areaABC float64
	plane   Plane
}

func NewTriangle(a, b, c Vector3, color Vector3) Triangle {
	plane := NewPlane(a, Subtract(b, a), Subtract(c, a))
	areaABC := ScalarProduct3(plane.n, CrossProduct(Subtract(b, a), Subtract(c, a)))
	return Triangle{
		a:       a,
		b:       b,
		c:       c,
		color:   color,
		areaABC: areaABC,
		plane:   plane,
	}
}

func NewTriangleWithName(a, b, c Vector3, color Vector3, name string) Triangle {
	result := NewTriangle(a, b, c, color)
	result.name = name
	return result
}

func (t Triangle) IsHit(ray *Ray) bool {
	if ray.HitsPlane(t.plane) {
		hitPoint, _ := ray.HitPoint()

		areaPBC := ScalarProduct3(t.plane.n, CrossProduct(Subtract(t.b, hitPoint), Subtract(t.c, hitPoint)))
		areaPCA := ScalarProduct3(t.plane.n, CrossProduct(Subtract(t.c, hitPoint), Subtract(t.a, hitPoint)))

		alpha := areaPBC / t.areaABC
		beta := areaPCA / t.areaABC
		gamma := 1.0 - alpha - beta

		if alpha >= 0.0 && alpha <= 1.0 && beta >= 0.0 && beta <= 1.0 && gamma >= 0.0 && gamma <= 1.0 {
			return true
		}
	}
	return false
}

func (t Triangle) String() string {
	return fmt.Sprintf("%s: A{%s}\nB{%s}\nC{%s}", t.name, t.A(), t.B(), t.C())
}

func (t Triangle) A() Vector3 {
	return t.a
}

func (t Triangle) B() Vector3 {
	return t.b
}

func (t Triangle) C() Vector3 {
	return t.c
}

func (t Triangle) Color() Vector3 {
	return t.color
}

func (t Triangle) Transform(matrix Matrix44) Triangle {
	return NewTriangle(
		NewVector4FromVector3(t.a, 1).Transform(matrix).ExtractVector3(),
		NewVector4FromVector3(t.b, 1).Transform(matrix).ExtractVector3(),
		NewVector4FromVector3(t.c, 1).Transform(matrix).ExtractVector3(),
		t.color,
	)
}

func (t *Triangle) DividePerspective(nearClippingPlane float64) {
	t.a = NewVector3((nearClippingPlane*t.a.x)/-t.a.z, (nearClippingPlane*t.a.y)/-t.a.z, -t.a.z)
	t.b = NewVector3((nearClippingPlane*t.b.x)/-t.b.z, (nearClippingPlane*t.b.y)/-t.b.z, -t.b.z)
	t.c = NewVector3((nearClippingPlane*t.c.x)/-t.c.z, (nearClippingPlane*t.c.y)/-t.c.z, -t.c.z)
}

func (t *Triangle) MapToNdcSpace(top, right, bottom, left float64) {
	t.a = NewVector3((2*t.a.x)/(right-left)-(right+left)/(right-left), (2*t.a.y)/(top-bottom)-(top+bottom)/(top-bottom), t.a.z)
	t.b = NewVector3((2*t.b.x)/(right-left)-(right+left)/(right-left), (2*t.b.y)/(top-bottom)-(top+bottom)/(top-bottom), t.b.z)
	t.c = NewVector3((2*t.c.x)/(right-left)-(right+left)/(right-left), (2*t.c.y)/(top-bottom)-(top+bottom)/(top-bottom), t.c.z)
}

func (t *Triangle) MapToRasterSpacer(height, width int) {
	t.a = NewVector3((t.a.x+1)/2*float64(width), (1-t.a.y)/2*float64(height), t.a.z)
	t.b = NewVector3((t.b.x+1)/2*float64(width), (1-t.b.y)/2*float64(height), t.b.z)
	t.c = NewVector3((t.c.x+1)/2*float64(width), (1-t.c.y)/2*float64(height), t.c.z)
}
