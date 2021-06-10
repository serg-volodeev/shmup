package bullet

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/serg-volodeev/shmup/internal/shape"
)

type Bullets struct {
	items   []*Bullet
	options *Opts
}

func NewBullets(o *Opts) *Bullets {
	b := &Bullets{}
	b.items = make([]*Bullet, 0, 20)
	b.options = o
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

func (b *Bullets) NewBullet(shipRect *shape.Rect) *Bullet {
	var item *Bullet

	for i := range b.items {
		if !b.items[i].visible {
			b.items[i].visible = true
			item = b.items[i]
			break
		}
	}

	if item == nil {
		item = newBullet(b.options)
		b.items = append(b.items, item)
	}

	item.rect.SetCenterX(shipRect.CenterX())
	item.rect.SetBottom(shipRect.Top())
	return item
}
