package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/serg-volodeev/shmup/internal/shape"
	"github.com/serg-volodeev/shmup/internal/shmup/res"
)

const (
	Title        = "Shmup"
	ScreenWidth  = 480
	ScreenHeight = 600
)

type Game struct {
	rect    *shape.Rect
	res     *res.Res
	space   *Space
	ship    *Ship
	bullets *Bullets
	meteors *Meteors
}

func NewGame() *Game {
	g := &Game{}
	g.res = res.LoadRes()
	g.rect = shape.NewRect(0, 0, ScreenWidth, ScreenHeight)
	g.space = newSpace(g)
	g.ship = newShip(g)
	g.bullets = newBullets(g)
	g.meteors = newMeteors(g)
	return g
}

func (g *Game) Update() error {
	g.bullets.update()
	g.meteors.update()
	return g.ship.update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.space.draw(screen)
	g.ship.draw(screen)
	g.bullets.draw(screen)
	g.meteors.draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
