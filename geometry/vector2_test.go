package geometry_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gnampfelix/grender/geometry"
)

var _ = Describe("Vector2", func() {
	It("should create a new Vector2", func() {
		a := NewVector2(10, 20)
		Expect(a.X()).Should(Equal(10.0))
		Expect(a.Y()).Should(Equal(20.0))
	})

	It("should add a Vector2", func() {
		a := NewVector2(10, 20)
		a.Add(a)
		Expect(a.X()).Should(Equal(20.0))
		Expect(a.Y()).Should(Equal(40.0))
	})

	It("should subtract a Vector2", func() {
		a := NewVector2(10, 20)
		a.Subtract(a)
		Expect(a.X()).Should(Equal(0.0))
		Expect(a.Y()).Should(Equal(0.0))
	})

	It("should calculate the length of the Vector2", func() {
		a := NewVector2(3, 0)
		l := a.Length()
		Expect(l).Should(Equal(3.0))
	})

	It("should multiply the Vector2 with a scalar", func() {
		a := NewVector2(10, 20)
		a.MultiplyWithScalar(3)
		Expect(a.X()).Should(Equal(30.0))
		Expect(a.Y()).Should(Equal(60.0))
	})

	It("should calculate the scalar product", func() {
		a := NewVector2(10, 20)
		scalarProduct := ScalarProduct2(a, a)
		Expect(scalarProduct).Should(Equal(500.0))
	})
	It("should calculate if the point is left of edge", func() {
		a := NewVector2(2.0, 1.0)
		b := NewVector2(3.0, 3.0)
		c := NewVector2(4.0, 1.0)
		p1 := NewVector2(3.0, 3.0)
		p2 := NewVector2(1.0, 1.0)

		Expect(CalculateEdgeFunction2(a, b, p1)).Should(BeNumerically(">=", 0.0))
		Expect(CalculateEdgeFunction2(b, c, p1)).Should(BeNumerically(">=", 0.0))
		Expect(CalculateEdgeFunction2(c, a, p1)).Should(BeNumerically(">=", 0.0))
		Expect(CalculateEdgeFunction2(a, b, p2)).Should(BeNumerically("<", 0.0))
	})
})
