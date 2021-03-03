package geometry

import (
	"fmt"
	"math"
)

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

// Add adds the Vector3 b to the vector by adding all of its components.
func (v *Vector3) Add(b Vector3) {
	v.x += b.x
	v.y += b.y
	v.z += b.z
}

// Subtract subtracts the vector b from the vector by subtracting all of its
// components
func (v *Vector3) Subtract(b Vector3) {
	v.x -= b.x
	v.y -= b.y
	v.z -= b.z
}

// MultiplyWithScalar multiplies the vector with the given scalar.
func (v *Vector3) MultiplyWithScalar(s float64) {
	v.x *= s
	v.y *= s
	v.z *= s
}

func (v Vector3) ExtractVector2() Vector2 {
	return NewVector2(v.x, v.y)
}

// Length calculates the length of the vector.
func (v Vector3) Length() float64 {
	sumOfSquares := v.x*v.x + v.y*v.y + v.z*v.z
	return math.Sqrt(sumOfSquares)
}

func (v *Vector3) Normalize() {
	norm := v.Length()
	v.MultiplyWithScalar(1.0 / norm)
}

func (v Vector3) String() string {
	return fmt.Sprintf("X: %f, Y: %f, Z: %f", v.x, v.y, v.z)
}

// ScalarProduct3 calculates the scalar (dot) product of two vector3.
func ScalarProduct3(a, b Vector3) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

// CrossProduct returns the cross product of two vectors.
func CrossProduct(a, b Vector3) Vector3 {
	newX := a.y*b.z - a.z*b.y
	newY := a.z*b.x - a.x*b.z
	newZ := a.x*b.y - a.y*b.x

	return NewVector3(newX, newY, newZ)
}

func AngleBetween(a, b Vector3) float64 {
	numerator := ScalarProduct3(a, b)
	denominator := a.Length() * b.Length()
	return math.Acos(numerator / denominator)
}

func Add(a, b Vector3) Vector3 {
	a.Add(b)
	return a
}

func Subtract(a, b Vector3) Vector3 {
	a.Subtract(b)
	return a
}
