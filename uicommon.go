package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math"
)

func DrawTriangle(screen *ebiten.Image, location Vector, velocity Vector, mainColor color.RGBA, secondaryColor color.RGBA) {
	intSize := 12
	size := float64(intSize)

	angle := math.Atan2(velocity.Y, velocity.X) + math.Pi/2
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	cx := location.X
	cy := location.Y

	vertices := make([]ebiten.Vertex, 3)

	ver1 := []float64{0, -size} // Top
	vertices[0] = creteVertex(ver1[0], ver1[1], sin, cos, cx, cy, secondaryColor)

	ver2 := []float64{-size * math.Sin(math.Pi/3), size / 2} // left
	vertices[1] = creteVertex(ver2[0], ver2[1], sin, cos, cx, cy, mainColor)

	ver3 := []float64{size * math.Sin(math.Pi/3), size / 2} // right
	vertices[2] = creteVertex(ver3[0], ver3[1], sin, cos, cx, cy, mainColor)

	indices := []uint16{0, 1, 2}

	whiteImg := ebiten.NewImage(1, 1)
	whiteImg.Fill(color.White)

	screen.DrawTriangles(vertices, indices, whiteImg, &ebiten.DrawTrianglesOptions{
		Filter:    ebiten.FilterNearest,
		AntiAlias: true,
	})
}

func creteVertex(lx float64, ly float64, sin float64, cos float64, cx float64, cy float64, color color.RGBA) ebiten.Vertex {
	x := lx*cos - ly*sin + cx
	y := lx*sin + ly*cos + cy

	return ebiten.Vertex{
		DstX: float32(x),
		DstY: float32(y),

		ColorA: float32(color.A) / 255.0,
		ColorR: float32(color.R) / 255.0,
		ColorG: float32(color.G) / 255.0,
		ColorB: float32(color.B) / 255.0,
	}
}

// mainly for debugging!!
func drawLine(screen *ebiten.Image, v1, v2 Vector) {
	dx := v2.X - v1.X
	dy := v2.Y - v2.Y
	length := math.Hypot(dx, dy)
	angle := math.Atan2(dy, dx)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(length, 1)
	op.GeoM.Rotate(angle)
	op.GeoM.Translate(v1.X, v2.Y)

	op.ColorScale.SetB(255)
	op.ColorScale.SetR(255)

	img := ebiten.NewImage(1, 1)

	screen.DrawImage(img, op)
}
