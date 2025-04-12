package main

import "github.com/hajimehoshi/ebiten/v2"

type Element interface {
	UpdateLocation()
	Draw(screen *ebiten.Image)
}
