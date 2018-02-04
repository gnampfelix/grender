package geometry_test

import (
	. "github.com/gnampfelix/grender/geometry"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Triangle", func() {
	It("should hit the triangle", func() {
		tri := NewTriangle(NewVector3(-1, -1, 0), NewVector3(1, -1, 0), NewVector3(0, 1 ,0), NewVector3(0, 0, 0))
		ray := NewRay(NewVector3(0, 0, 1), NewVector3(0, 0, -1))

		isHit := tri.IsHit(&ray)
		Expect(isHit).Should(Equal(true))
	})
})
