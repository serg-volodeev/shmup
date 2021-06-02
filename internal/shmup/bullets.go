package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Bullets struct {
	res   *Res
	items []*Bullet
}

func newBullets(res *Res) *Bullets {
	b := &Bullets{}
	b.res = res
	b.items = make([]*Bullet, 0, 20)
	return b
}

func (b *Bullets) update(world *World) {
	for i := range b.items {
		b.items[i].update(world)
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

	item := newBullet(b.res)
	b.items = append(b.items, item)
	return item
}
