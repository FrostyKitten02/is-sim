package main

import "math/rand"

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
