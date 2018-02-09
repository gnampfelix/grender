package renderer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gnampfelix/grender/renderer"
)

var _ = Describe("Depthbuffer", func() {
	It("should set a value", func() {
		buffer := NewMapBuffer()
		set := buffer.SetDepthIfCloser(12.0, 0, 0)
		Expect(set).Should(BeTrue())
	})

	It("should override a value", func() {
		buffer := NewMapBuffer()
		set := buffer.SetDepthIfCloser(12.0, 0, 0)
		Expect(set).Should(BeTrue())

		set = buffer.SetDepthIfCloser(10.0, 0, 0)
		Expect(set).Should(BeTrue())
		})

		It("should not override a value", func() {
			buffer := NewMapBuffer()
			set := buffer.SetDepthIfCloser(12.0, 0, 0)
			Expect(set).Should(BeTrue())

			set = buffer.SetDepthIfCloser(14.0, 0, 0)
			Expect(set).Should(BeFalse())
			})
})
