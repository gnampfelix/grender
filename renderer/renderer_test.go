package renderer_test

import (
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"

	"github.com/gnampfelix/grender/geometry"
	. "github.com/gnampfelix/grender/renderer"
	"strconv"
)

var _ = Describe("Renderer", func() {
	It("should render", func() {
		object := NewCube()
		input := NewInput()
		input.Add(object)
		renderer := New()

		for i := 0; i < 1; i++ {
			object.Rotate(geometry.Z, 1.2)
			output := renderer.Render(input)
			output.Save("tmp/"+strconv.Itoa(i)+".png")
		}
	})
})
