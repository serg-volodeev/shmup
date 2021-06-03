package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	rect    *Rect
	space   *Space
	ship    *Ship
	bullets *Bullets
	meteors *Meteors
}

func newWorld(res *Res) *World {
	w := &World{}
	w.rect = newRect(0, 0, ScreenWidth, ScreenHeight)
	w.space = newSpace(res)
	w.ship = newShip(res, w)
	w.bullets = newBullets(res)
	w.meteors = newMeteors(res, w)
	return w
}

func (w *World) update() error {
	if err := w.ship.update(w); err != nil {
		return err
	}
	w.bullets.update(w)
	w.meteors.update(w)
	return nil
}

func (w *World) draw(screen *ebiten.Image) {
	w.space.draw(screen)
	w.ship.draw(screen)
	w.bullets.draw(screen)
	w.meteors.draw(screen)
}
