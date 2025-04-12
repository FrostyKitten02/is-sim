package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math"
)

type Vector struct {
	X float32
	Y float32
}

type Agent struct {
	Location  *Location
	Direction *Vector //how much speed in used for each coordinate, then this is how much we move the agent for (x*speed, y*speed)
}

// TODO move!
func GetVector(l1 Location, l2 Location) Vector {
	x := l2.X - l1.X
	y := l2.Y - l1.Y

	return Vector{x, y}
}

func GetVecLen(v Vector) float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func NormalizeVector(v Vector) Vector {
	vecLen := float32(GetVecLen(v))
	if vecLen == 0 {
		return Vector{
			0, 0,
		}
	}
	return Vector{
		X: v.X / vecLen,
		Y: v.Y / vecLen,
	}
}

func SumVec(v1 Vector, v2 Vector) Vector {
	return Vector{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

func ReversVecDirection(v Vector) Vector {
	return Vector{
		X: -v.X,
		Y: -v.Y,
	}
}

func ScaleVec(v Vector, scale float32) Vector {
	return Vector{
		X: v.X * scale,
		Y: v.Y * scale,
	}
}

// TODO check why when updating location and direction on agent values don't get updated out of scope, why are they not reflected in draw call
func (a *Agent) UpdateLocation(gs *GameState) {
	//calculating optimal direction to target
	optimalDir := GetVector(*a.Location, *gs.target)
	optimalDirNorm := NormalizeVector(optimalDir)

	//calculating needed force/forceVec to change direction of agent to match optimalDirection
	reverseCurrDir := ReversVecDirection(*a.Direction)
	reverseCurrDirNorm := NormalizeVector(reverseCurrDir)
	forceVec := SumVec(optimalDirNorm, reverseCurrDirNorm)
	forceVecNorm := NormalizeVector(forceVec)
	neededForce := GetVecLen(forceVecNorm)

	//making checks for maxForce bsc agents have max rotation (maxForce)
	if neededForce <= gs.maxForce {
		//TODO: maybe combine direction with forceVec?
		a.Direction.X = optimalDirNorm.X
		a.Direction.Y = optimalDirNorm.Y
	} else {
		scaleTo := gs.maxForce / neededForce
		scaledForceVec := ScaleVec(forceVecNorm, float32(scaleTo))
		newDirectionVec := SumVec(*a.Direction, scaledForceVec)
		newDirectionVecNorm := NormalizeVector(newDirectionVec)

		a.Direction.X = newDirectionVecNorm.X
		a.Direction.Y = newDirectionVecNorm.Y
	}

	//lastly update agent position using new direction
	xSpeed := a.Direction.X * gs.maxSpeed
	ySpeed := a.Direction.Y * gs.maxSpeed
	a.Location.X = a.Location.X + xSpeed
	a.Location.Y = a.Location.Y + ySpeed
}

func (a *Agent) Draw(screen *ebiten.Image) {
	size := float32(10)

	angle := math.Atan2(float64(a.Direction.Y), float64(a.Direction.X)) + math.Pi/2
	cos := float32(math.Cos(angle))
	sin := float32(math.Sin(angle))
	cx := a.Location.X
	cy := a.Location.Y

	local := []Location{
		{0, -size}, // Top
		{-size * float32(math.Sin(math.Pi/3)), size / 2}, // left
		{size * float32(math.Sin(math.Pi/3)), size / 2},  // right
	}

	vertices := make([]ebiten.Vertex, 3)
	for i, pt := range local {
		lx, ly := pt.X, pt.Y
		x := lx*cos - ly*sin + cx
		y := lx*sin + ly*cos + cy

		vertices[i] = ebiten.Vertex{
			DstX:   x,
			DstY:   y,
			ColorR: 1, ColorG: 1, ColorB: 0, ColorA: 1,
		}
	}

	indices := []uint16{0, 1, 2}

	whiteImg := ebiten.NewImage(1, 1)
	whiteImg.Fill(color.White)

	screen.DrawTriangles(vertices, indices, whiteImg, &ebiten.DrawTrianglesOptions{
		Filter:    ebiten.FilterNearest,
		AntiAlias: true,
	})
}
