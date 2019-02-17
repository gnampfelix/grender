package main

import (
	"encoding/json"
	"github.com/gnampfelix/pub"
	"github.com/gnampfelix/grender/geometry"
	"github.com/gnampfelix/grender/renderer"
)

var myPub pub.Publisher

type streamerOutput struct {
		height, width int
		image         []Line
		depthBuffer   []float64
}

func NewStreamerOutput(height, width int) renderer.Output {
	image := make([]Line, height)
	for i := 0; i < height; i++{
		line := make([]Dot, width)
		image[i] = line
	}
	return &streamerOutput{
		height: height,
		width:  width,
		image:  image,
	}
}

func (s streamerOutput)Height() int {
	return s.height;
}

func (s streamerOutput)Width() int {
	return s.width;
}

func (s *streamerOutput)SetPixel (c geometry.Vector3, x, y int) {
	s.image[y][x] = Dot{
		R: int(c.X()),
		B: int(c.Y()),
		G: int(c.Z()),
	}
	if x == len(s.image[y]) {
		message := pub.NewMessage("ws")
		enc := json.NewEncoder(message)
		enc.Encode(s.image[y])
		myPub.Publish(message)
	}
}

func (s streamerOutput)Save(a string){}

type Dot struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

type Line []Dot
