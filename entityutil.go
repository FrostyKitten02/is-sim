package main

func seek(gs *GameState, target Vector, location Vector, velocity Vector) Vector {
	desired := SubVectors(target, location)
	desiredLimited := MagVec(desired, gs.maxSpeed)

	//steer
	steer := SubVectors(desiredLimited, velocity)
	steerLimited := LimitVec(steer, gs.maxForce)
	return steerLimited
}

func wrapBorders(gs *GameState, location *Vector) {
	if location.X < -gs.wanderR {
		location.X = gs.width + gs.wanderR
	}

	if location.Y < -gs.wanderR {
		location.Y = gs.height + gs.wanderR
	}

	if location.X > gs.width+gs.wanderR {
		location.X = -gs.wanderR
	}

	if location.Y > gs.height+gs.wanderR {
		location.Y = -gs.wanderR
	}
}
