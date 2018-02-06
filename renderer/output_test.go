package renderer_test

import (
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"

	. "github.com/gnampfelix/grender/renderer"
	"github.com/gnampfelix/grender/geometry"
)

var _ = Describe("Output", func() {
	It("should create an image", func() {
		output := NewSimpleOutput(20, 20)
		for i := 0; i < 20; i++ {
			for j := 0; j < 20; j++ {
				output.SetPixel(geometry.NewVector3(10, float64(i*i), float64(j*j)), i, j)
			}
		}
		output.Save()
	})
})
