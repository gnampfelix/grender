package renderer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gnampfelix/grender/renderer"
)

var _ = Describe("Object", func() {
	It("should return some tris", func() {
		input := NewCube()

		Expect(input.HasNextTriangle()).Should(BeTrue())
		nextTri := input.NextTriangle()
		Expect(nextTri.Color().X()).Should(Equal(255.0))
		Expect(nextTri.A().Z()).Should(Equal(-5.0))

		Expect(input.HasNextTriangle()).Should(BeTrue())
		nextTri = input.NextTriangle()
		Expect(nextTri.B().X()).Should(Equal(5.0))
		Expect(nextTri.C().Z()).Should(Equal(5.0))

		Expect(input.HasNextTriangle()).Should(BeTrue())
		nextTri = input.NextTriangle()
		Expect(nextTri.C().Y()).Should(Equal(20.0))

		Expect(input.HasNextTriangle()).Should(BeTrue())
		nextTri = input.NextTriangle()
		Expect(nextTri.A().Z()).Should(Equal(-5.0))

		nextTri = input.NextTriangle()
		Expect(nextTri.Color().Y()).Should(Equal(255.0))
	})
})
