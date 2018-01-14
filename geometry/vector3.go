package geometry

// A Vector3 represents a vector in the three dimensional space.
type Vector3 struct {
	x, y, z float64
}

// NewVector3 creates a new Vector3 instance.
func NewVector3(x, y, z float64) Vector3 {
	return Vector3{
		x: x,
		y: y,
		z: z,
	}
}

// X returns the x component of the vector3.
func (v Vector3) X() float64 {
	return v.x
}

// Y returns the y component of the vector3.
func (v Vector3) Y() float64 {
	return v.y
}

// Z returns the z component of the vector3.
func (v Vector3) Z() float64 {
	return v.z
}
