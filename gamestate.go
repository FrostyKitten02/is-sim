package main

type GameState struct {
	Agents   *[]Agent
	Elements *[]Element
	target   *Circle

	maxSpeed       float64
	maxForce       float64
	arriveDistance float64
	playAreaOffset int
	width          int
	height         int
}

func (gs *GameState) InitGameState(width int, height int) {
	gs.maxSpeed = 8
	gs.maxForce = 0.05
	gs.arriveDistance = 200
	gs.playAreaOffset = 100
	gs.width = width
	gs.height = height

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
