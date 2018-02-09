package geometry

import (
	"fmt"
)

type Triangle struct {
	a, b, c Vector3
	color   Vector3
}

func NewTriangle(a, b, c Vector3, color Vector3) Triangle {
	return Triangle{
		a:     a,
		b:     b,
		c:     c,
		color: color,
	}
}

func (t Triangle) IsHit(ray *Ray) bool {
	plane := NewPlane(t.a, Subtract(t.b, t.a), Subtract(t.c, t.a))
	if ray.HitsPlane(plane) {
		hitPoint, _ := ray.HitPoint()

		areaABC := ScalarProduct3(plane.Normal(), CrossProduct(Subtract(t.b, t.a), Subtract(t.c, t.a)))
		areaPBC := ScalarProduct3(plane.Normal(), CrossProduct(Subtract(t.b, hitPoint), Subtract(t.c, hitPoint)))
		areaPCA := ScalarProduct3(plane.Normal(), CrossProduct(Subtract(t.c, hitPoint), Subtract(t.a, hitPoint)))

		alpha := areaPBC / areaABC
		beta := areaPCA / areaABC
		gamma := 1.0 - alpha - beta

		if alpha >= 0.0 && alpha <= 1.0 && beta >= 0.0 && beta <= 1.0 && gamma >= 0.0 && gamma <= 1.0 {
			return true
		}
	}
	return false
}

func (t Triangle)String()string{
	return fmt.Sprintf("A{%s}\nB{%s}\nC{%s}", t.A(), t.B(), t.C())
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
