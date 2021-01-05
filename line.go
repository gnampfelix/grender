package main

import (
	"encoding/json"
	"github.com/gnampfelix/grender/geometry"
	"github.com/gnampfelix/grender/renderer"
	"github.com/gnampfelix/pub"
)

var myPub pub.Publisher

type streamerOutput struct {
	height, width int
}

func NewStreamerOutput(height, width int) renderer.Output {
	return &streamerOutput{
		height: height,
		width:  width,
	}
}

func (s streamerOutput) Height() int {
	return s.height
}

func (s streamerOutput) Width() int {
	return s.width
}

func (s *streamerOutput) SetPixel(c geometry.Vector3, x, y int) {
	message := pub.NewMessage("ws")
	enc := json.NewEncoder(message)
	enc.Encode(Dot{
		X: x,
		Y: y,
		R: int(c.X()),
		G: int(c.Y()),
		B: int(c.Z()),
	})
	myPub.Publish(message)
}

func (s streamerOutput) Save(a string) {}

type Dot struct {
	X int `json:"x"`
	Y int `json:"y"`
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}
