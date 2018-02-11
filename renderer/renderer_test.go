package renderer_test

import (
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"

	. "github.com/gnampfelix/grender/renderer"
	"github.com/gnampfelix/grender/geometry"
)

var _ = Describe("Renderer", func() {
	It("should render", func() {
		object := NewCube()
		input := NewInput()
		transformation := geometry.NewMatrix44(
			geometry.NewVector4(0.707, -0.707, 0, 0),
			geometry.NewVector4(0.707, 0.707, 0, 0),
			geometry.NewVector4(0, 0, 1, 0),
			geometry.NewVector4(0, 0, 0, 1),
		)
		object.Transform(transformation)
		input.Add(object)
		renderer := New()
		output := renderer.Render(input)
		output.Save("renderer.png")
	})
})
