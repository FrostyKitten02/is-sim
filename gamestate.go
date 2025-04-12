package main

type Location struct {
	X float32
	Y float32
}

type GameState struct {
	Agents   *[]Agent
	Elements *[]Element
	target   *Location

	maxSpeed float32
	maxForce float64
}

func (gs *GameState) InitGameState() {
	gs.maxSpeed = 10
	//should be between 0 - sqrt(2)
	gs.maxForce = 0.2

	targetCircle := &Circle{
		&Location{
			X: 800,
			Y: 300,
		},
	}
	gs.target = targetCircle.Location

	gs.Elements = &[]Element{
		targetCircle,
	}

	gs.Agents = &[]Agent{
		//{
		//	Location: &Location{
		//		X: 40,
		//		Y: 40,
		//	},
		//	Direction: &Vector{
		//		X: 0,
		//		Y: 0,
		//	},
		//},
		{
			Location: &Location{
				X: 100,
				Y: 500,
			},
			Direction: &Vector{
				X: -1,
				Y: -1,
			},
		},
		//{
		//	Location: &Location{
		//		X: 0,
		//		Y: 0,
		//	},
		//	Direction: &Vector{
		//		X: 0,
		//		Y: 0,
		//	},
		//},
	}
}
