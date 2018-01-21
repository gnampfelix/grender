package geometry

// A Plane represents a plane in 3D space:
// x = p + s * u + t * v where x, p, u, v are element of R³.
// OR
// 0 = (x - p) * n where x, p, n are emelent of R³
// OR
// n1 * x1 + n2 *x2 + n3 * x3 = d where (n1, n2, n3) = n & (x1, x2, x3) = x
type Plane struct {
	p, u, v, n Vector3
	d          float64
}

// NewPlane creates a new Plane in three dimensional space that is spanned by u
// and v around p.
func NewPlane(p, u, v Vector3) Plane {
	result := Plane{
		p: p,
		u: u,
		v: v,
	}
	result.calculateNormal()
	result.u.Normalize()
	result.v.Normalize()
	result.d = ScalarProduct3(result.n, result.p)
	return result
}

// P returns the location vector of the plane.
func (p Plane) P() Vector3 {
	return p.p
}

// U returns one normalized direction vector of the plane.
func (p Plane) U() Vector3 {
	return p.u
}

// V returns another normalized location vector of the plane.
func (p Plane) V() Vector3 {
	return p.v
}

// Normal returns the normal vector of the plane.
func (p Plane) Normal() Vector3 {
	return p.n
}

// D returns the value of the d component so that
// n1 * x1 + n2 *x2 + n3 * x3 = d where (n1, n2, n3) = n & (x1, x2, x3) = x
// is true for all x that are on the plane.
func (p Plane) D() float64 {
	return p.d
}

func (p *Plane) calculateNormal() {
	p.n = CrossProduct(p.u, p.v)
	p.n.Normalize()
}
