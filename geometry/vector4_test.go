package geometry_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gnampfelix/grender/geometry"
)

var _ = Describe("Vector4", func() {
	It("should Create a new Vector4", func() {
		a := NewVector4FromVector3(NewVector3(1, 2, 3), 4)
		Expect(a.A()).Should(Equal(4.0))
	})

	It("should Create a new Matrix44", func() {
		m := NewMatrix44(
			NewVector4(1, 2, 3, 4),
			NewVector4(1, 2, 3, 4),
			NewVector4(1, 2, 3, 4),
			NewVector4(1, 2, 3, 4),
		)

		Expect(m.D().A()).Should(Equal(4.0))
	})

	It("should transform a Vector4", func() {
		v := NewVector4(0, 1, 0, 1)
		m := NewMatrix44(
			NewVector4(0, 0, 1, 1.5),
			NewVector4(0, -1, 0, 1),
			NewVector4(1, 0, 0, 1.5),
			NewVector4(0, 0, 0, 1),
		)

		v2 := v.Transform(m)
		Expect(v2.X()).Should(Equal(1.5))
		Expect(v2.Y()).Should(Equal(0.0))
		Expect(v2.Z()).Should(Equal(1.5))
		Expect(v2.A()).Should(Equal(1.0))
	})
})
