package renderer_test

import (
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"

	"github.com/gnampfelix/grender/geometry"
	. "github.com/gnampfelix/grender/renderer"
)

var _ = Describe("Renderer", func() {
	It("should render", func() {
		object := NewCube()
		input := NewInput()

		object.Rotate(geometry.Z, 45)
		input.Add(object)
		renderer := New()
		output := renderer.Render(input)
		output.Save("renderer.png")
	})
})
