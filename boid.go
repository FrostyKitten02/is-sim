package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Forces struct {
	SeparationForce Vector
	AlignForce      Vector
	CohereForce     Vector
}

type Boid struct {
	Id           uint
	Location     *Vector
	Acceleration *Vector
	Velocity     *Vector
	WanderTheta  *float64 //not used anymore?
}

var boidMainColor = color.RGBA{
	A: 120,
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
	forces := a.calcForces(gs)

	a.ApplyForce(ScaleVec(forces.SeparationForce, 1.5))
	a.ApplyForce(ScaleVec(forces.AlignForce, 1))
	a.ApplyForce(ScaleVec(forces.CohereForce, 1))

	a.update(gs)
	a.wrapBorders(gs)
}

func (a *Boid) update(gs *GameState) {
	//updating values on agent
	*a.Velocity = LimitVec(SumVec(*a.Velocity, *a.Acceleration), gs.maxSpeed)
	*a.Location = SumVec(*a.Location, *a.Velocity)

	a.Acceleration.X = 0
	a.Acceleration.Y = 0
}

func (a *Boid) seek(gs *GameState, target Vector) Vector {
	return seek(gs, target, *a.Location, *a.Velocity)
}

func (a *Boid) calcForces(gs *GameState) Forces {
	forces := Forces{}
	//separation
	separationDist := gs.separationR * 2
	separationSum := Vector{
		X: 0,
		Y: 0,
	}
	separationCount := 0

	//align
	velocitySum := Vector{
		X: 0,
		Y: 0,
	}
	alignCount := 0

	//cohere
	positionSum := Vector{
		X: 0,
		Y: 0,
	}
	cohereCount := 0

	for _, boid := range gs.Flock.Boids {
		//do we always skip?? should align skip self?
		if boid.Id == a.Id {
			continue
		}

		//separation
		//important to sub vectors this - other, so we get inverse direction force
		diff := SubVectors(*a.Location, *boid.Location)
		dist := GetVecLen(diff)
		if dist <= separationDist {
			//inversely proportional to distance
			vec := MagVec(diff, 1/separationDist)
			separationSum = SumVec(separationSum, vec)
			separationCount++
		}

		//align
		if dist <= gs.alignDistance {
			velocitySum = SumVec(velocitySum, *boid.Velocity)
			alignCount++
		}

		//cohere
		if dist <= gs.cohereDistance {
			positionSum = SumVec(positionSum, *boid.Location)
			cohereCount++
		}

	}

	//separation
	if separationCount == 0 {
		forces.SeparationForce = separationSum
	} else {
		//limit and avg our separationSum vec
		avg := ScaleVec(separationSum, 1.0/float64(separationCount))
		limited := MagVec(avg, gs.maxSpeed)
		separation := SubVectors(limited, *a.Velocity)
		separationForce := LimitVec(separation, gs.maxForce)
		forces.SeparationForce = separationForce
	}

	//align
	if alignCount == 0 {
		forces.AlignForce = Vector{
			X: 0,
			Y: 0,
		}
	} else {
		avg := ScaleVec(velocitySum, 1.0/float64(alignCount))
		limited := MagVec(avg, gs.maxSpeed)
		alignSteer := SubVectors(limited, *a.Velocity)
		limitedAlignSteer := LimitVec(alignSteer, gs.maxForce)
		forces.AlignForce = limitedAlignSteer
	}

	//cohere
	if cohereCount == 0 {
		forces.CohereForce = Vector{
			X: 0,
			Y: 0,
		}
	} else {
		avg := ScaleVec(positionSum, 1.0/float64(cohereCount))
		forces.CohereForce = a.seek(gs, avg)
	}

	return forces
}

func (a *Boid) wrapBorders(gs *GameState) {
	wrapBorders(gs, a.Location)
}

func (a *Boid) ApplyForce(force Vector) {
	*a.Acceleration = SumVec(*a.Acceleration, force)
}

func (a *Boid) Draw(screen *ebiten.Image) {
	DrawTriangle(screen, *a.Location, *a.Velocity, boidMainColor, boidSecondColor)
}
