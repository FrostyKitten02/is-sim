package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct {
	state *GameState
}

func (g *Game) Update() error {
	for _, agent := range *g.state.Agents {
		agent.UpdateLocation(g.state)
	}

	for _, element := range *g.state.Elements {
		element.UpdateLocation(g.state)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, agent := range *g.state.Agents {
		agent.Draw(screen)
	}

	for _, element := range *g.state.Elements {
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

	ebiten.SetTPS(30)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
