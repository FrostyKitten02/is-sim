package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Agent struct {
	Location     *Vector
	Acceleration *Vector
	Velocity     *Vector
}

var agentMainColor = color.RGBA{
	A: 120,
	R: 255,
	G: 0,
	B: 0,
}

var agentSecondColor = color.RGBA{
	A: 255,
	R: 255,
	G: 0,
	B: 0,
}

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

	//area limit
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
	//in other implementations we use MagVec but for arrive Limit is better
	//it makes it so our agent doesn't turn in circles
	steerLimited := LimitVec(steer, gs.maxForce)
	a.ApplyForce(steerLimited)

	//updating values on agent
	*a.Velocity = LimitVec(SumVec(*a.Velocity, *a.Acceleration), gs.maxSpeed)
	*a.Location = SumVec(*a.Location, *a.Velocity)

	a.Acceleration.X = 0
	a.Acceleration.Y = 0
}

func (a *Agent) ApplyForce(force Vector) {
	*a.Acceleration = SumVec(*a.Acceleration, force)
}

func (a *Agent) Draw(screen *ebiten.Image) {
	DrawTriangle(screen, *a.Location, *a.Velocity, agentMainColor, agentSecondColor)
}
