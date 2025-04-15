package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math"
)

func creteVertex(lx float32, ly float32, sin float32, cos float32, cx float32, cy float32, color color.RGBA) ebiten.Vertex {
	x := lx*cos - ly*sin + cx
	y := lx*sin + ly*cos + cy

	return ebiten.Vertex{
		DstX: x,
		DstY: y,

		ColorA: float32(color.A / 255),
		ColorR: float32(color.R / 255),
		ColorG: float32(color.G / 255),
		ColorB: float32(color.B / 255),
	}
}

func DrawTriangle(screen *ebiten.Image, location Vector, velocity Vector, mainColor color.RGBA, secondaryColor color.RGBA) {
	size := float32(12)

	angle := math.Atan2(float64(velocity.Y), float64(velocity.X)) + math.Pi/2
	cos := float32(math.Cos(angle))
	sin := float32(math.Sin(angle))
	cx := location.X
	cy := location.Y

	vertices := make([]ebiten.Vertex, 3)

	ver1 := Vector{0, -size} // Top
	vertices[0] = creteVertex(ver1.X, ver1.Y, sin, cos, cx, cy, secondaryColor)

	ver2 := Vector{-size * float32(math.Sin(math.Pi/3)), size / 2} // left
	vertices[1] = creteVertex(ver2.X, ver2.Y, sin, cos, cx, cy, mainColor)

	ver3 := Vector{size * float32(math.Sin(math.Pi/3)), size / 2} // right
	vertices[2] = creteVertex(ver3.X, ver3.Y, sin, cos, cx, cy, mainColor)

	indices := []uint16{0, 1, 2}

	whiteImg := ebiten.NewImage(1, 1)
	whiteImg.Fill(color.White)

	screen.DrawTriangles(vertices, indices, whiteImg, &ebiten.DrawTrianglesOptions{
		Filter:    ebiten.FilterNearest,
		AntiAlias: true,
	})
}
