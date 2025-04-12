package main

type Location struct {
	X float32
	Y float32
}

type GameState struct {
	Agents   *[]Agent
	Elements *[]Element
	target   *Circle

	maxSpeed float32
	maxForce float64
}

func (gs *GameState) InitGameState() {
	gs.maxSpeed = 8
	//should be between 0 - sqrt(2)
	gs.maxForce = 0.2

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
}
