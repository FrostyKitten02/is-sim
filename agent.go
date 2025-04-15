package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math"
)

type Agent struct {
	Location     *Vector
	Acceleration *Vector
	Velocity     *Vector
}

var mainColor = color.RGBA{
	A: 0, //making invisible tail
	R: 120,
	G: 0,
	B: 0,
}

var secondColor = color.RGBA{
	A: 255,
	R: 255,
	G: 0,
	B: 0,
}

// TODO check why when updating location and direction on agent values don't get updated out of scope, why are they not reflected in draw call
func (a *Agent) UpdateLocation(gs *GameState) {
	desired := SubVectors(*gs.target.Location, *a.Location)
	distance := GetVecLen(desired)

	var desiredLimited Vector
	//arrive
	if distance < gs.arriveDistance {
		percent := distance / gs.arriveDistance
		desiredLimited = MagVec(desired, percent*gs.maxSpeed)
	} else {
		desiredLimited = MagVec(desired, gs.maxSpeed)
	}

	if a.Location.X < float32(gs.playAreaOffset) {
		desiredLimited = MagVec(Vector{X: float32(gs.maxSpeed), Y: a.Velocity.Y}, gs.maxSpeed)
	} else if a.Location.X > float32(gs.width-gs.playAreaOffset) {
		desiredLimited = MagVec(Vector{X: -float32(gs.maxSpeed), Y: a.Velocity.Y}, gs.maxSpeed)
	}

	if a.Location.Y < float32(gs.playAreaOffset) {
		desiredLimited = MagVec(Vector{X: a.Velocity.X, Y: float32(gs.maxSpeed)}, gs.maxSpeed)
	} else if a.Location.Y > float32(gs.height-gs.playAreaOffset) {
		desiredLimited = MagVec(Vector{X: a.Velocity.X, Y: -float32(gs.maxSpeed)}, gs.maxSpeed)
	}

	//steer
	steer := SubVectors(desiredLimited, *a.Velocity)
	steerLimited := LimitVec(steer, gs.maxForce)
	a.ApplyForce(steerLimited)

	//updating values on agent
	newVelocity := LimitVec(SumVec(*a.Velocity, *a.Acceleration), gs.maxSpeed)
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
	size := float32(12)

	angle := math.Atan2(float64(a.Velocity.Y), float64(a.Velocity.X)) + math.Pi/2
	cos := float32(math.Cos(angle))
	sin := float32(math.Sin(angle))
	cx := a.Location.X
	cy := a.Location.Y

	vertices := make([]ebiten.Vertex, 3)

	ver1 := Vector{0, -size} // Top
	vertices[0] = creteVertex(ver1.X, ver1.Y, sin, cos, cx, cy, secondColor)

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
