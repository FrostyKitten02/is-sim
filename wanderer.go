package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math"
)

type Wanderer struct {
	Location     *Vector
	Acceleration *Vector
	Velocity     *Vector
	WanderTheta  *float64
}

var wandererMainColor = color.RGBA{
	A: 120,
	R: 200,
	G: 0,
	B: 200,
}

var wandererSecondColor = color.RGBA{
	A: 255,
	R: 200,
	G: 0,
	B: 200,
}

func (a *Wanderer) UpdateLocation(gs *GameState) {
	target := a.wander(gs)

	steerForce := seek(gs, target, *a.Location, *a.Velocity)
	a.ApplyForce(steerForce)

	//updating values on agent
	newVelocity := LimitVec(SumVec(*a.Velocity, *a.Acceleration), gs.maxSpeed)
	a.Velocity.X = newVelocity.X
	a.Velocity.Y = newVelocity.Y

	newPosition := SumVec(*a.Location, *a.Velocity)
	a.Location.X = newPosition.X
	a.Location.Y = newPosition.Y

	a.Acceleration.X = 0
	a.Acceleration.Y = 0

	wrapBorders(gs, a.Location)
}

func (a *Wanderer) wander(gs *GameState) Vector {
	change := 0.5

	*a.WanderTheta = *a.WanderTheta + randomFloat(-change, change)
	circlePos := MagVec(*a.Velocity, gs.wanderD)
	circlePos = SumVec(circlePos, *a.Location)

	directionAngle := VecAngle(*a.Velocity)
	//creating offset vector for circle
	newDirection := Vector{
		X: gs.wanderR * math.Cos(*a.WanderTheta+directionAngle),
		Y: gs.wanderR * math.Sin(*a.WanderTheta+directionAngle),
	}

	return SumVec(circlePos, newDirection)
}

func (a *Wanderer) ApplyForce(force Vector) {
	updated := SumVec(*a.Acceleration, force)
	a.Acceleration.X = updated.X
	a.Acceleration.Y = updated.Y
}

func (a *Wanderer) Draw(screen *ebiten.Image) {
	DrawTriangle(screen, *a.Location, *a.Velocity, wandererMainColor, wandererSecondColor)
}
