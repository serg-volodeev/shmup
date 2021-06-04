package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/serg-volodeev/shmup/internal/shape"
)

type World struct {
	rect    *shape.Rect
	space   *Space
	ship    *Ship
	bullets *Bullets
	meteors *Meteors
}

func newWorld(res *Res) *World {
	w := &World{}
	w.rect = shape.NewRect(0, 0, ScreenWidth, ScreenHeight)
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
