package geometry

type Matrix44 struct {
	a, b, c, d Vector4
}

func NewMatrix44(a, b, c, d Vector4) Matrix44 {
	return Matrix44{
		a: a,
		b: b,
		c: c,
		d: d,
	}
}

func (m Matrix44) A() Vector4 {
	return m.a
}

func (m Matrix44) B() Vector4 {
	return m.b
}

func (m Matrix44) C() Vector4 {
	return m.c
}

func (m Matrix44) D() Vector4 {
	return m.d
}
