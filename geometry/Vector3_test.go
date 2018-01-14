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
})
