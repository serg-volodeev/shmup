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
	world *World
}

func NewGame() *Game {
	res := loadRes()
	g := &Game{}
	g.world = newWorld(res)
	return g
}

func (g *Game) Update() error {
	return g.world.update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
