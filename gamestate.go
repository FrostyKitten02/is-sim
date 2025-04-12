package main

type Location struct {
	X float32
	Y float32
}

type GameState struct {
	agents   []Agent
	elements []Element
}

func (gs *GameState) InitGameState() {
	gs.elements = []Element{
		&Circle{
			Location{
				X: 800,
				Y: 300,
			},
		},
	}

	gs.agents = []Agent{
		{
			Location: &Location{
				X: 40,
				Y: 40,
			},
			SpeedVec: &SpeedVector{
				X: 1,
				Y: 0,
			},
		},
		{
			Location: &Location{
				X: 0,
				Y: 0,
			},
			SpeedVec: &SpeedVector{
				X: 0,
				Y: 1,
			},
		},
		{
			Location: &Location{
				X: 0,
				Y: 0,
			},
			SpeedVec: &SpeedVector{
				X: 1,
				Y: 1,
			},
		},
	}
}
