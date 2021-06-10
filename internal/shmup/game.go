package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/serg-volodeev/shmup/internal/shape"
	"github.com/serg-volodeev/shmup/internal/shmup/bullet"
	"github.com/serg-volodeev/shmup/internal/shmup/meteor"
	"github.com/serg-volodeev/shmup/internal/shmup/res"
	"github.com/serg-volodeev/shmup/internal/shmup/ship"
	"github.com/serg-volodeev/shmup/internal/shmup/space"
)

const (
	Title        = "Shmup"
	ScreenWidth  = 480
	ScreenHeight = 600
)

type Game struct {
	rect    *shape.Rect
	res     *res.Res
	space   *space.Space
	ship    *ship.Ship
	bullets *bullet.Bullets
	meteors *meteor.Meteors
}

func NewGame() *Game {
	g := &Game{}
	g.res = res.LoadRes()
	g.rect = shape.NewRect(0, 0, ScreenWidth, ScreenHeight)
	g.space = space.NewSpace(&space.Opts{Res: g.res})
	g.meteors = meteor.NewMeteors(&meteor.Opts{Res: g.res, Bounds: g.rect})
	g.bullets = bullet.NewBullets(&bullet.Opts{Res: g.res, Bounds: g.rect, Meteors: g.meteors})
	g.ship = ship.NewShip(&ship.Opts{Res: g.res, Bounds: g.rect, Meteors: g.meteors, Bullets: g.bullets})
	return g
}

func (g *Game) Update() error {
	g.bullets.Update()
	g.meteors.Update()
	return g.ship.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.space.Draw(screen)
	g.ship.Draw(screen)
	g.bullets.Draw(screen)
	g.meteors.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
