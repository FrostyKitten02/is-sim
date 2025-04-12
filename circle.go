package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Circle struct {
	Location *Vector
}

func (p *Circle) UpdateLocation(gs *GameState) {
	//NOT MOVING!
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		p.Location.X = float32(x)
		p.Location.Y = float32(y)
	}
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
