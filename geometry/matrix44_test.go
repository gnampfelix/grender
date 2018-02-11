package geometry_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gnampfelix/grender/geometry"
)

var _ = Describe("Matrix44", func() {
	It("should chain two matrices", func() {
		a := NewMatrix44(
			NewVector4(1, 2, 3, 4),
			NewVector4(5, 6, 7, 8),
			NewVector4(9, 10, 11, 12),
			NewVector4(13, 14, 15, 16),
		)
		b := NewMatrix44(
			NewVector4(1, 0, 1, 0),
			NewVector4(0, 1, 0, 1),
			NewVector4(1, 0, 1, 0),
			NewVector4(0, 1, 0, 1),
		)

		result := a.Chain(b)

		Expect(result.A().X()).Should(Equal(4.0))
		Expect(result.B().Y()).Should(Equal(14.0))
		Expect(result.C().Z()).Should(Equal(20.0))
		Expect(result.D().A()).Should(Equal(30.0))
	})
})
