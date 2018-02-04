package geometry_test

import (
	. "github.com/gnampfelix/grender/geometry"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ray", func() {
	It("should hit the plane", func() {
		ray := NewRay(NewVector3(0, 0, -1), NewVector3(0, 0, 1))
		plane := NewPlane(NewVector3(0,0,0), NewVector3(1,0,0), NewVector3(0,1,0))

		isHit := ray.HitsPlane(plane)
		Expect(isHit).Should(Equal(true))
	})
})
