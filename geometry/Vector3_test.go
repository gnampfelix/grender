package geometry_test

import (
	. "github.com/gnampfelix/grender/geometry"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vector3", func() {
	It("should create a new Vector3", func() {
		a := NewVector3(10, 20, 30)
		Expect(a.X()).Should(Equal(10.0))
		Expect(a.Y()).Should(Equal(20.0))
		Expect(a.Z()).Should(Equal(30.0))
	})

	It("should add a Vector3", func() {
		a := NewVector3(10, 20, 30)
		a.Add(a)
		Expect(a.X()).Should(Equal(20.0))
		Expect(a.Y()).Should(Equal(40.0))
		Expect(a.Z()).Should(Equal(60.0))
	})

	It("should subtract a Vector3", func() {
		a := NewVector3(10, 20, 30)
		a.Subtract(a)
		Expect(a.X()).Should(Equal(0.0))
		Expect(a.Y()).Should(Equal(0.0))
		Expect(a.Z()).Should(Equal(0.0))
	})

	It("should calculate the length of the Vector3", func() {
		a := NewVector3(3, 0, 4)
		l := a.Length()
		Expect(l).Should(Equal(5.0))
	})

	It("should calculate the cross product of two vectors", func() {
		a := NewVector3(1, 2, 3)
		b := NewVector3(3, 2, 1)
		c := CrossProduct(a, b)
		Expect(c.X()).Should(Equal(-4.0))
		Expect(c.Y()).Should(Equal(8.0))
		Expect(c.Z()).Should(Equal(-4.0))
	})

	It("should multiply the Vector3 with a scalar", func() {
		a := NewVector3(10, 20, 30)
		a.MultiplyWithScalar(3)
		Expect(a.X()).Should(Equal(30.0))
		Expect(a.Y()).Should(Equal(60.0))
		Expect(a.Z()).Should(Equal(90.0))
	})

	It("should calculate the scalar product", func() {
		a := NewVector3(10, 20, 30)
		scalarProduct := ScalarProduct3(a, a)
		Expect(scalarProduct).Should(Equal(1400.0))
	})

	It("should calculate the angle between two Vector3", func() {
		a := NewVector3(3, 2, 1)
		b := NewVector3(1, 2, 3)
		angle := AngleBetween(a, b)
		Expect(angle).Should(BeNumerically("~", 0.77519337331036130720409371118247))
	})

	It("should normalize a vector", func() {
		a := NewVector3(3, 0, 4)
		a.Normalize()
		Expect(a.X()).Should(BeNumerically("~", 3.0/5.0))
		Expect(a.Y()).Should(BeNumerically("~", 0.0/5.0))
		Expect(a.Z()).Should(BeNumerically("~", 4.0/5.0))
	})

	It("should subtract two vectors", func() {
		a := NewVector3(1,2,3)
		b := NewVector3(1,2,3)
		c := Subtract(a, b)
		Expect(c.X()).Should(Equal(0.0))
		Expect(a.X()).Should(Equal(1.0))
	})
})
