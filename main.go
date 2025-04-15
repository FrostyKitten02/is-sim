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

	for _, agent := range *g.state.Agents {
		agent.UpdateLocation(g.state)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	for _, element := range *g.state.Elements {
		element.Draw(screen)
	}

	for _, agent := range *g.state.Agents {
		agent.Draw(screen)
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

	ebiten.SetTPS(60)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
