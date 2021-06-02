package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	rect    *Rect
	space   *Space
	ship    *Ship
	bullets *Bullets
}

func newWorld(res *Res) *World {
	w := &World{}
	w.rect = newRect(0, 0, ScreenWidth, ScreenHeight)
	w.space = newSpace(res)
	w.ship = newShip(res, w)
	w.bullets = newBullets(res)
	return w
}

func (w *World) update() {
	w.ship.update(w)
	w.bullets.update(w)
}

func (w *World) draw(screen *ebiten.Image) {
	w.space.draw(screen)
	w.ship.draw(screen)
	w.bullets.draw(screen)
}
