package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math"
)

type SpeedVector struct {
	X float32
	Y float32
}

type Agent struct {
	Location *Location
	SpeedVec *SpeedVector
}

func (a *Agent) UpdateLocation() {
	a.Location.X = a.Location.X + a.SpeedVec.X
	a.Location.Y = a.Location.Y + a.SpeedVec.Y
}

func (a *Agent) Draw(screen *ebiten.Image) {
	size := float32(10)

	angle := math.Atan2(float64(a.SpeedVec.Y), float64(a.SpeedVec.X)) + math.Pi/2
	cos := float32(math.Cos(angle))
	sin := float32(math.Sin(angle))
	cx := a.Location.X
	cy := a.Location.Y

	local := []Location{
		{0, -size}, // Top
		{-size * float32(math.Sin(math.Pi/3)), size / 2}, // left
		{size * float32(math.Sin(math.Pi/3)), size / 2},  // right
	}

	vertices := make([]ebiten.Vertex, 3)
	for i, pt := range local {
		lx, ly := pt.X, pt.Y
		x := lx*cos - ly*sin + cx
		y := lx*sin + ly*cos + cy

		vertices[i] = ebiten.Vertex{
			DstX:   x,
			DstY:   y,
			ColorR: 1, ColorG: 1, ColorB: 0, ColorA: 1,
		}
	}

	indices := []uint16{0, 1, 2}

	whiteImg := ebiten.NewImage(1, 1)
	whiteImg.Fill(color.White)

	screen.DrawTriangles(vertices, indices, whiteImg, &ebiten.DrawTrianglesOptions{
		Filter:    ebiten.FilterNearest,
		AntiAlias: true,
	})
}
