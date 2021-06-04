package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Bullets struct {
	game  *Game
	items []*Bullet
}

func newBullets(g *Game) *Bullets {
	b := &Bullets{}
	b.game = g
	b.items = make([]*Bullet, 0, 20)
	return b
}

func (b *Bullets) update() {
	for i := range b.items {
		b.items[i].update()
	}
}

func (b *Bullets) draw(screen *ebiten.Image) {
	for i := range b.items {
		b.items[i].draw(screen)
	}
}

func (b *Bullets) newBullet() *Bullet {
	for i := range b.items {
		if !b.items[i].visible {
			b.items[i].visible = true
			return b.items[i]
		}
	}

	item := newBullet(b.game)
	b.items = append(b.items, item)
	return item
}
