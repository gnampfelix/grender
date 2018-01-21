package geometry_test

import (
	. "github.com/gnampfelix/grender/geometry"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Plane", func() {
	It("should create a plane", func() {
		p := NewVector3(0, 0, 0)
		u := NewVector3(2, 0, 0)
		v := NewVector3(0, 2, 0)
		plane := NewPlane(p, u, v)
		Expect(plane.U().X()).Should(Equal(1.0))
	})

	It("should have the correct normal vector", func() {
		p := NewVector3(0, 0, 0)
		u := NewVector3(2, 0, 0)
		v := NewVector3(0, 2, 0)
		plane := NewPlane(p, u, v)
		Expect(plane.Normal().X()).Should(Equal(0.0))
		Expect(plane.Normal().Y()).Should(Equal(0.0))
		Expect(plane.Normal().Z()).Should(Equal(1.0))
	})
})
