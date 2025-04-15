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
	alignDistance  float64
	cohereDistance float64
}

func (gs *GameState) InitGameState(width int, height int) {
	gs.maxSpeed = 3
	gs.maxForce = 0.05
	gs.arriveDistance = 200
	gs.playAreaOffset = 100
	gs.width = float64(width)
	gs.height = float64(height)
	gs.wanderR = 3
	gs.wanderD = 80
	gs.separationR = 25
	gs.alignDistance = 50
	gs.cohereDistance = 50

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

	boidLen := 20
	boids := make([]Boid, boidLen)
	for i := 0; i < boidLen; i++ {
		theta := 0.3 + float64(i)/float64(boidLen)
		boids[i] = Boid{
			Id: uint(i),
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
	}
	gs.Flock = &Flock{
		boids,
	}
}
