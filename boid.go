package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math"
)

type Boid struct {
	Location     *Vector
	Acceleration *Vector
	Velocity     *Vector
	WanderTheta  *float64
}

var boidMainColor = color.RGBA{
	A: 0, //making invisible tail
	R: 0,
	G: 0,
	B: 120,
}

var boidSecondColor = color.RGBA{
	A: 255,
	R: 0,
	G: 0,
	B: 120,
}

// TODO check why when updating location and direction on agent values don't get updated out of scope, why are they not reflected in draw call
func (a *Boid) UpdateLocation(gs *GameState) {
	target := a.wander()
	a.wrapBorders(gs)
	a.update(gs)
	a.seek(gs, target)
}

func (a *Boid) update(gs *GameState) {
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

// TODO pass in target!
func (a *Boid) seek(gs *GameState, target Vector) {
	desired := SubVectors(target, *a.Location)
	desiredLimited := MagVec(desired, gs.maxSpeed)

	//steer
	steer := SubVectors(desiredLimited, *a.Velocity)
	steerLimited := LimitVec(steer, gs.maxForce)
	a.ApplyForce(steerLimited)
}

func (a *Boid) wrapBorders(gs *GameState) {
	if a.Location.X < -gs.wanderR {
		a.Location.X = gs.width + gs.wanderR
	}

	if a.Location.Y < -gs.wanderR {
		a.Location.Y = gs.height + gs.wanderR
	}

	if a.Location.X > gs.width+gs.wanderR {
		a.Location.X = -gs.wanderR
	}

	if a.Location.Y > gs.height+gs.wanderR {
		a.Location.Y = -gs.wanderR
	}
}

func (a *Boid) wander() Vector {
	wanderR := float64(25)
	wanderD := float64(80)
	change := 0.3

	*a.WanderTheta = *a.WanderTheta + randomFloat(-change, change)
	circlePos := MagVec(*a.Velocity, wanderD)
	circlePos = SumVec(circlePos, *a.Location)

	angleOffset := VecAngle(*a.Velocity)
	//creating offset vector for circle
	circleOffset := Vector{
		X: wanderR * math.Cos(*a.WanderTheta+angleOffset),
		Y: wanderR * math.Sin(*a.WanderTheta+angleOffset),
	}

	return SumVec(circlePos, circleOffset)
}

func (a *Boid) ApplyForce(force Vector) {
	updated := SumVec(*a.Acceleration, force)
	a.Acceleration.X = updated.X
	a.Acceleration.Y = updated.Y
}

func (a *Boid) Draw(screen *ebiten.Image) {
	DrawTriangle(screen, *a.Location, *a.Velocity, boidMainColor, boidSecondColor)
}
