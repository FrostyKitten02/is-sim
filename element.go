package main

import "github.com/hajimehoshi/ebiten/v2"

type Element interface {
	UpdateLocation(state *GameState)
	Draw(screen *ebiten.Image)
}
