package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type World struct {
	space   *Space
	ship    *Ship
	bullets *Bullets
}

func newWorld(res *Res) *World {
	w := &World{}
	w.space = newSpace(res)
	w.ship = newShip(res)
	w.bullets = newBullets(res)
	return w
}

func (w *World) update() {
	w.ship.update()
	w.bullets.update()

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		w.fire()
	}
}

func (w *World) draw(screen *ebiten.Image) {
	w.space.draw(screen)
	w.ship.draw(screen)
	w.bullets.draw(screen)
}

func (w *World) fire() {
	b := w.bullets.newBullet()
	b.x = w.ship.x + w.ship.width()/2 - b.width()/2
	b.y = ScreenHeight - w.ship.height() - 10 - b.height()
}
