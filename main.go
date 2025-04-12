package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	state *GameState
}

func (g *Game) Update() error {
	for _, agent := range g.state.agents {
		agent.UpdateLocation()
	}

	for _, element := range g.state.elements {
		element.UpdateLocation()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, agent := range g.state.agents {
		agent.Draw(screen)
	}

	for _, element := range g.state.elements {
		element.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("IS-sim!")
	game := &Game{
		state: &GameState{},
	}
	game.state.InitGameState()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
