package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Circle struct {
	Location Location
}

func (p *Circle) UpdateLocation() {
	//NOT MOVING!
}

func (p *Circle) Draw(screen *ebiten.Image) {
	pointColor := color.RGBA{
		A: 0,
		R: 255,
		G: 0,
		B: 0,
	}
	vector.DrawFilledCircle(screen, p.Location.X, p.Location.Y, 10, pointColor, true)
}
