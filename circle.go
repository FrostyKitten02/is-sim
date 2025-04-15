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
		p.Location.X = float64(x)
		p.Location.Y = float64(y)
	}
}

func (p *Circle) Draw(screen *ebiten.Image) {
	pointColor := color.RGBA{
		A: 230,
		R: 0,
		G: 255,
		B: 0,
	}
	vector.DrawFilledCircle(screen, float32(p.Location.X), float32(p.Location.Y), 10, pointColor, true)
}
