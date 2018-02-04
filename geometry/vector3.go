package geometry

import (
	"math"
	"fmt"
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

// Length calculates the length of the vector.
func (v Vector3) Length() float64 {
	sumOfSquares := math.Pow(v.X(), 2) + math.Pow(v.Y(), 2) + math.Pow(v.Z(), 2)
	return math.Sqrt(sumOfSquares)
}

func (v *Vector3) Normalize() {
	norm := v.Length()
	v.MultiplyWithScalar(1.0 / norm)
}

func (v Vector3)String() string {
	return fmt.Sprintf("X: %f, Y: %f, Z: %f", v.x, v.y, v.z)
}

// ScalarProduct3 calculates the scalar (dot) product of two vector3.
func ScalarProduct3(a, b Vector3) float64 {
	result := a.X() * b.X()
	result += a.Y() * b.Y()
	result += a.Z() * b.Z()
	return result
}

// CrossProduct returns the cross product of two vectors.
func CrossProduct(a, b Vector3) Vector3 {
	newX := a.Y()*b.Z() - a.Z()*b.Y()
	newY := a.Z()*b.X() - a.X()*b.Z()
	newZ := a.X()*b.Y() - a.Y()*b.X()

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
