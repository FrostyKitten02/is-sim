package main

import "image/color"

type Flock struct {
	Boids       *[]Boid
	Name        string
	mainColor   color.RGBA
	secondColor color.RGBA
}

type GameState struct {
	//all elements
	Elements *[]Element
	//flock sim
	Wanderer *Wanderer

	Flocks map[string]*Flock
	//follow and land
	Agents *[]Agent
	target *Circle

	maxSpeed       float64
	maxForce       float64
	arriveDistance float64
	playAreaOffset float64
	width          float64
	height         float64
	wanderR        float64
	wanderD        float64
	separationR    float64
	alignDistance  float64
	cohereDistance float64
}

func (gs *GameState) InitGameState(width int, height int) {
	gs.maxSpeed = 8
	gs.maxForce = 0.1
	gs.arriveDistance = 200
	gs.playAreaOffset = 100
	gs.width = float64(width)
	gs.height = float64(height)
	gs.wanderR = 3
	gs.wanderD = 80
	gs.separationR = 25
	gs.alignDistance = 70
	gs.cohereDistance = 70
	gs.Flocks = map[string]*Flock{}

	targetCircle := &Circle{
		&Vector{
			X: 800,
			Y: 300,
		},
	}
	gs.target = targetCircle

	gs.Elements = &[]Element{
		targetCircle,
	}

	gs.Agents = &[]Agent{
		{
			Location: &Vector{
				X: 100,
				Y: 500,
			},
			Acceleration: &Vector{
				X: 0,
				Y: 0,
			},
			Velocity: &Vector{
				X: 0,
				Y: 0,
			},
		},
	}

	//adding agents to elements
	for _, agent := range *gs.Agents {
		*gs.Elements = append(*gs.Elements, &agent)
	}

	boidMainColor1 := color.RGBA{
		A: 120,
		R: 0,
		G: 0,
		B: 120,
	}

	boidSecondColor1 := color.RGBA{
		A: 255,
		R: 0,
		G: 0,
		B: 120,
	}

	//TODO should think about modifying repel force, so flock are seperated...
	//boidMainColor2 := color.RGBA{
	//	A: 120,
	//	R: 120,
	//	G: 120,
	//	B: 0,
	//}
	//
	//boidSecondColor2 := color.RGBA{
	//	A: 255,
	//	R: 120,
	//	G: 120,
	//	B: 0,
	//}

	createFlock(gs, "FLOCK-1", boidMainColor1, boidSecondColor1)
	//createFlock(gs, "FLOCK-2", boidMainColor2, boidSecondColor2)

	wandererTheta := 0.3 + randomFloat(-0.3, 0.7)
	gs.Wanderer = &Wanderer{
		Location: &Vector{
			X: 600,
			Y: 600,
		},
		Acceleration: &Vector{
			X: 0,
			Y: 0,
		},
		Velocity: &Vector{
			X: -1,
			Y: 0.2,
		},
		WanderTheta: &wandererTheta,
	}

	el := append(*gs.Elements, gs.Wanderer)
	*gs.Elements = el
}

func createFlock(gs *GameState, name string, mainColor, secondColor color.RGBA) {
	boidLen := 100
	boids := make([]Boid, boidLen)
	for i := 0; i < boidLen; i++ {
		theta := 0.3 + float64(i)/float64(boidLen)
		boid := &Boid{
			Id:          uint(i),
			FlockName:   name,
			mainColor:   mainColor,
			secondColor: secondColor,
			Location: &Vector{
				X: 400 + float64(i) + randomFloat(float64(boidLen), 70),
				Y: 400 + float64(i) + randomFloat(float64(boidLen), 70),
			},
			WanderTheta: &theta,
			Acceleration: &Vector{
				X: 0,
				Y: 0,
			},
			Velocity: &Vector{
				X: randomFloat(-2, 2),
				Y: randomFloat(-2, 2),
			},
		}
		boids[i] = *boid

		//adding boids to elements
		*gs.Elements = append(*gs.Elements, boid)
	}

	gs.Flocks[name] = &Flock{
		Boids: &boids,
		Name:  name,
	}
}
