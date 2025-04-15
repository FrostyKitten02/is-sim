package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Boid struct {
	Location     *Vector
	Acceleration *Vector
	Velocity     *Vector
}

var boidMainColor = color.RGBA{
	A: 0, //making invisible tail
	R: 120,
	G: 0,
	B: 0,
}

var boidSecondColor = color.RGBA{
	A: 255,
	R: 255,
	G: 0,
	B: 0,
}

// TODO check why when updating location and direction on agent values don't get updated out of scope, why are they not reflected in draw call
func (a *Boid) UpdateLocation(gs *GameState) {
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

	if a.Location.X < gs.playAreaOffset {
		desiredLimited = MagVec(Vector{X: gs.maxSpeed, Y: a.Velocity.Y}, gs.maxSpeed)
	} else if a.Location.X > gs.width-gs.playAreaOffset {
		desiredLimited = MagVec(Vector{X: -gs.maxSpeed, Y: a.Velocity.Y}, gs.maxSpeed)
	}

	if a.Location.Y < gs.playAreaOffset {
		desiredLimited = MagVec(Vector{X: a.Velocity.X, Y: gs.maxSpeed}, gs.maxSpeed)
	} else if a.Location.Y > gs.height-gs.playAreaOffset {
		desiredLimited = MagVec(Vector{X: a.Velocity.X, Y: -gs.maxSpeed}, gs.maxSpeed)
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

func (a *Boid) ApplyForce(force Vector) {
	updated := SumVec(*a.Acceleration, force)
	a.Acceleration.X = updated.X
	a.Acceleration.Y = updated.Y
}

func (a *Boid) Draw(screen *ebiten.Image) {
	DrawTriangle(screen, *a.Location, *a.Velocity, boidMainColor, boidSecondColor)
}
