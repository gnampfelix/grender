package renderer

import (
	"github.com/gnampfelix/grender/geometry"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

type Output interface {
	Height() int
	Width() int
	SetPixel(c geometry.Vector3, x, y int)
	Save(filename string)
}

func NewSimpleOutput(height, width int) Output {
	image := image.NewRGBA(image.Rect(0, 0, width, height))
	return &simpleOutput{
		height: height,
		width:  width,
		image:  image,
	}
}

type simpleOutput struct {
	height, width int
	image         *image.RGBA
	depthBuffer   []float64
}

func (s simpleOutput) Height() int {
	return s.height
}

func (s simpleOutput) Width() int {
	return s.width
}

func (s *simpleOutput) SetPixel(c geometry.Vector3, x, y int) {
	s.image.SetRGBA(x, y, color.RGBA{uint8(c.X()), uint8(c.Y()), uint8(c.Z()), 255})
}

func (s simpleOutput) Save(filename string) {
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		return
	}
	_ = jpeg.Encode(f, s.image, nil)

}
