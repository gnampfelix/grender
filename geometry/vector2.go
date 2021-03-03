package geometry

import (
	"fmt"
	"math"
)

// A Vector2 represents a vector in the three dimensional space.
type Vector2 struct {
	x, y float64
}

// NewVector2 creates a new Vector3 instance.
func NewVector2(x, y float64) Vector2 {
	return Vector2{
		x: x,
		y: y,
	}
}

// X returns the x component of the vector2.
func (v Vector2) X() float64 {
	return v.x
}

// Y returns the y component of the vector2.
func (v Vector2) Y() float64 {
	return v.y
}

// Add adds the Vector3 b to the vector by adding all of its components.
func (v *Vector2) Add(b Vector2) {
	v.x += b.x
	v.y += b.y
}

// Subtract subtracts the vector b from the vector by subtracting all of its
// components
func (v *Vector2) Subtract(b Vector2) {
	v.x -= b.x
	v.y -= b.y
}

// MultiplyWithScalar multiplies the vector with the given scalar.
func (v *Vector2) MultiplyWithScalar(s float64) {
	v.x *= s
	v.y *= s
}

// Length calculates the length of the vector.
func (v Vector2) Length() float64 {
	sumOfSquares := v.x*v.x + v.y*v.y
	return math.Sqrt(sumOfSquares)
}

func (v *Vector2) Normalize() {
	norm := v.Length()
	v.MultiplyWithScalar(1.0 / norm)
}

func (v Vector2) String() string {
	return fmt.Sprintf("X: %f, Y: %f", v.x, v.y)
}

// ScalarProduct2 calculates the scalar (dot) product of two vector3.
func ScalarProduct2(a, b Vector2) float64 {
	return a.x*b.x + a.y*b.y
}

// CalculateEdgeFunction2 returns a value > 0
// if the point lies on the right side of the edge described by a and b
// and < 0 otherwise. If the value is 0, the point is exactly on the edge.
func CalculateEdgeFunction2(a, b Vector2, point Vector2) float64 {
	return (point.x-a.x)*(b.y-a.y) - (point.y-a.y)*(b.x-a.x)
}
