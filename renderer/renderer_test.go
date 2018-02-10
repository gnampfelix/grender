package renderer_test

import (
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"

	. "github.com/gnampfelix/grender/renderer"
)

var _ = Describe("Renderer", func() {
	It("should render", func() {
		input := NewCubeInput()
		renderer := New()
		output := renderer.Render(input)
		output.Save("renderer.png")
	})
})
