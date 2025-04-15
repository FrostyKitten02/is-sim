package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

type Game struct {
	state *GameState
}

func (g *Game) Update() error {
	for _, element := range *g.state.Elements {
		element.UpdateLocation(g.state)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	for _, element := range *g.state.Elements {
		element.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(g.state.width), int(g.state.height)
}

func main() {
	width := 1280
	height := 720

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("IS-sim!")
	game := &Game{
		state: &GameState{},
	}
	game.state.InitGameState(width, height)

	ebiten.SetTPS(60)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
