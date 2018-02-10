package geometry

type Vector4 struct {
	x, y, z, a float64
}

func NewVector4FromVector3(base Vector3, a float64) Vector4 {
	return Vector4{
		x: base.x,
		y: base.y,
		z: base.z,
		a: a,
	}
}

func NewVector4(x, y, z, a float64) Vector4 {
	return Vector4{
		x: x,
		y: y,
		z: z,
		a: a,
	}
}

func (v Vector4) ExtractVector3() Vector3 {
	return NewVector3(v.x, v.y, v.z)
}

func (v Vector4) X() float64 {
	return v.x
}

func (v Vector4) Y() float64 {
	return v.y
}

func (v Vector4) Z() float64 {
	return v.z
}

func (v Vector4) A() float64 {
	return v.a
}

func ScalarProduct4(a, b Vector4) float64 {
	result := a.X() * b.X()
	result += a.Y() * b.Y()
	result += a.Z() * b.Z()
	result += a.A() * b.A()
	return result
}

func (v Vector4) Transform(matrix Matrix44) Vector4{
	newX := ScalarProduct4(matrix.A(), v)
	newY := ScalarProduct4(matrix.B(), v)
	newZ := ScalarProduct4(matrix.C(), v)
	newA := ScalarProduct4(matrix.D(), v)

	return NewVector4(newX, newY, newZ, newA)
}
