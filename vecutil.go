package main

import "math"

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

func SubVectors(v1 Vector, v2 Vector) Vector {
	v2Reverse := ReversVecDirection(v2)
	return SumVec(v1, v2Reverse)
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

func LimitVec(v Vector, limit float64) Vector {
	length := GetVecLen(v)
	if length < limit {
		return v
	}

	normalized := NormalizeVector(v)
	return ScaleVec(normalized, float32(limit))
}

func MagVec(v Vector, limit float64) Vector {
	length := GetVecLen(v)
	if length == limit {
		return v
	}

	normalized := NormalizeVector(v)
	return ScaleVec(normalized, float32(limit))
}
