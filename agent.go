package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math"
)

type Vector struct {
	X float32
	Y float32
}

type Agent struct {
	Location     *Vector
	Acceleration *Vector
	Velocity     *Vector
}

// TODO check why when updating location and direction on agent values don't get updated out of scope, why are they not reflected in draw call
func (a *Agent) UpdateLocation(gs *GameState) {
	desired := SubVectors(*gs.target, *a.Location)
	desiredLimited := MagVec(desired, float64(gs.maxSpeed))

	steer := SubVectors(desiredLimited, *a.Velocity)
	steerLimited := LimitVec(steer, gs.maxForce)
	a.ApplyForce(steerLimited)

	//TODO update values!! for position!!
	newVelocity := LimitVec(SumVec(*a.Velocity, *a.Acceleration), float64(gs.maxSpeed))
	a.Velocity.X = newVelocity.X
	a.Velocity.Y = newVelocity.Y

	newPosition := SumVec(*a.Location, *a.Velocity)
	a.Location.X = newPosition.X
	a.Location.Y = newPosition.Y

	a.Acceleration.X = 0
	a.Acceleration.Y = 0
}

func (a *Agent) ApplyForce(force Vector) {
	updated := SumVec(*a.Acceleration, force)
	a.Acceleration.X = updated.X
	a.Acceleration.Y = updated.Y
}

func (a *Agent) Draw(screen *ebiten.Image) {
	size := float32(10)

	angle := math.Atan2(float64(a.Velocity.Y), float64(a.Velocity.X)) + math.Pi/2
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
