package main

type Location struct {
	X float32
	Y float32
}

type GameState struct {
	agents   []Agent
	elements []Element
}
