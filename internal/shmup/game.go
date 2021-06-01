package shmup

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Title        = "Shmup"
	ScreenWidth  = 320
	ScreenHeight = 240
)

type Game struct {
	ship *Ship
}

func NewGame() *Game {
	game := &Game{}
	game.ship = NewShip()
	return game
}

func (g *Game) Update() error {
	// Write your game's logical update.
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	screen.Fill(color.White)
	g.ship.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
