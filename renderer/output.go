package renderer

import (
  "github.com/gnampfelix/grender/geometry"
  "image/color"
  "image"
  "image/jpeg"
  "os"
)

type Output interface {
  Height() int
  Width() int
  SetPixel(c geometry.Vector3, x, y int)
  SetDepth(depth float64, x, y int)
  DepthAt(x, y int) float64
  Save()
}

func NewSimpleOutput(height, width int) Output{
  depth := make([]float64, height*width)
  image := image.NewRGBA(image.Rect(0, 0, width, height))
  return &simpleOutput{
    height: height,
    width: width,
    image: image,
    depthBuffer: depth,
  }
}

type simpleOutput struct {
  height, width int
  image *image.RGBA
  depthBuffer []float64
}

func (s simpleOutput)Height()int{
  return s.height
}

func (s simpleOutput)Width()int{
  return s.width
}

func (s *simpleOutput)SetPixel(c geometry.Vector3, x, y int) {
  s.image.SetRGBA(x, y, color.RGBA{uint8(c.X()), uint8(c.Y()), uint8(c.Z()), 255})
}

func (s *simpleOutput)SetDepth(depth float64, x, y int) {
  if x < s.width && y < s.height {
    s.depthBuffer[x*y] = depth
  }
}

func (s simpleOutput)DepthAt(x, y int) float64 {
  if  x < s.width && y < s.height {
    return s.depthBuffer[x*y]
  }
  return 0
}

func (s simpleOutput)Save() {
  f, err := os.Create("image.jpg")
  defer f.Close()
  if err != nil {
    return
  }
  _ = jpeg.Encode(f, s.image, nil)

}
