package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Title        = "Shmup"
	ScreenWidth  = 480
	ScreenHeight = 600
)

type Game struct {
	space *Space
	ship  *Ship
}

func NewGame() *Game {
	g := &Game{}
	g.space = NewSpace()
	g.ship = NewShip()
	return g
}

func (g *Game) Update() error {
	g.ship.Update()
	updateBullets()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.space.Draw(screen)
	g.ship.Draw(screen)
	drawBullets(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
