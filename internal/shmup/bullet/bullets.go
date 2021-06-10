package bullet

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/serg-volodeev/shmup/internal/shape"
)

type Bullets struct {
	items []*Bullet
	opts  *Opts
}

func NewBullets(o *Opts) *Bullets {
	b := &Bullets{}
	b.items = make([]*Bullet, 0, 20)
	b.opts = o
	return b
}

func (b *Bullets) Update() {
	for i := range b.items {
		b.items[i].update()
	}
}

func (b *Bullets) Draw(screen *ebiten.Image) {
	for i := range b.items {
		b.items[i].draw(screen)
	}
}

func (b *Bullets) findFree() *Bullet {
	for i := range b.items {
		if !b.items[i].visible {
			b.items[i].visible = true
			return b.items[i]
		}
	}
	return nil
}

func (b *Bullets) NewBullet(shipRect *shape.Rect) *Bullet {
	item := b.findFree()

	if item == nil {
		item = newBullet(b.opts)
		b.items = append(b.items, item)
	}

	item.setPos(shipRect)
	return item
}
