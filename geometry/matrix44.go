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

func (m Matrix44) Chain(matrix Matrix44) Matrix44 {
	col1 := NewVector4(matrix.a.x, matrix.b.x, matrix.c.x, matrix.d.x)
	col2 := NewVector4(matrix.a.y, matrix.b.y, matrix.c.y, matrix.d.y)
	col3 := NewVector4(matrix.a.z, matrix.b.z, matrix.c.z, matrix.d.z)
	col4 := NewVector4(matrix.a.a, matrix.b.a, matrix.c.a, matrix.d.a)

	a := NewVector4(
		ScalarProduct4(m.a, col1),
		ScalarProduct4(m.a, col2),
		ScalarProduct4(m.a, col3),
		ScalarProduct4(m.a, col4),
	)

	b := NewVector4(
		ScalarProduct4(m.b, col1),
		ScalarProduct4(m.b, col2),
		ScalarProduct4(m.b, col3),
		ScalarProduct4(m.b, col4),
	)

	c := NewVector4(
		ScalarProduct4(m.c, col1),
		ScalarProduct4(m.c, col2),
		ScalarProduct4(m.c, col3),
		ScalarProduct4(m.c, col4),
	)

	d := NewVector4(
		ScalarProduct4(m.d, col1),
		ScalarProduct4(m.d, col2),
		ScalarProduct4(m.d, col3),
		ScalarProduct4(m.d, col4),
	)
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
