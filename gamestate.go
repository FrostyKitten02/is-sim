package main

type Flock struct {
	Boids []Boid
}

type GameState struct {
	Agents   *[]Agent
	Elements *[]Element
	target   *Circle
	Flock    *Flock

	maxSpeed       float64
	maxForce       float64
	arriveDistance float64
	playAreaOffset float64
	width          float64
	height         float64
	wanderR        float64
	wanderD        float64
	separationR    float64
}

func (gs *GameState) InitGameState(width int, height int) {
	gs.maxSpeed = 2
	gs.maxForce = 0.05
	gs.arriveDistance = 200
	gs.playAreaOffset = 100
	gs.width = float64(width)
	gs.height = float64(height)
	gs.wanderR = 20
	gs.wanderD = 80
	gs.separationR = 12

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

	theta1 := 0.0
	theta2 := 0.75
	gs.Flock = &Flock{
		[]Boid{
			{
				Id: 1,
				Location: &Vector{
					X: 400,
					Y: 400,
				},
				WanderTheta: &theta1,
				Acceleration: &Vector{
					X: 0,
					Y: 0,
				},
				Velocity: &Vector{
					X: 1,
					Y: 0,
				},
			},
			{
				Id: 2,
				Location: &Vector{
					X: 400,
					Y: 410,
				},
				WanderTheta: &theta2,
				Acceleration: &Vector{
					X: 0,
					Y: 0,
				},
				Velocity: &Vector{
					X: 1,
					Y: 0,
				},
			},
		},
	}
}
