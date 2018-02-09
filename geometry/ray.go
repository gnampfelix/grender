package geometry

import "errors"

// Ray represents a ray in three dimensional space.
type Ray struct {
	p, u, hitPoint Vector3
	isHit          bool
}

func NewRay(p, u Vector3) Ray {
	ray := Ray{
		p: p,
		u: u,
	}
	ray.u.Normalize()
	return ray
}

func (r Ray) P() Vector3 {
	return r.p
}

func (r Ray) U() Vector3 {
	return r.u
}

func (r *Ray) HitsPlane(p Plane) bool {
	normal := p.Normal()
	// If the normal of the plane and the direction of the ray are orthogonal
	// (scalarProduct = 0), ray and plane are parallel.
	denominator := ScalarProduct3(r.u, normal)
	if denominator > 1e-6 || denominator < -(1e-6) { //not parallel or within
		tmp := p.P()
		tmp.Subtract(r.p)
		t := ScalarProduct3(tmp, normal) / denominator
		if t >= 0 {
			hit := NewVector3(r.u.X(), r.u.Y(), r.u.Z())
			hit.MultiplyWithScalar(t)
			hit.Add(r.p)
			r.hitPoint = hit
			r.isHit = true
			return true
		}
	}
	return false
}

func (r Ray) HitPoint() (Vector3, error) {
	if r.isHit {
		return r.hitPoint, nil
	}
	return Vector3{}, errors.New("ray does not hit anything")
}
