package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	rect    *Rect
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
	speedY  float64
	visible bool
	radius  float64
}

func newBullet(res *Res) *Bullet {
	b := &Bullet{}
	b.image = res.images["bullet"]
	b.options = &ebiten.DrawImageOptions{}
	b.speedY = -8
	b.visible = true
	b.rect = newRectFromImage(b.image)
	b.radius = b.rect.w / 2
	return b
}

func (b *Bullet) draw(screen *ebiten.Image) {
	if !b.visible {
		return
	}
	b.options.GeoM.Reset()
	b.options.GeoM.Translate(b.rect.x, b.rect.y)
	screen.DrawImage(b.image, b.options)
}

func (b *Bullet) update(world *World) {
	if !b.visible {
		return
	}
	b.rect.moveY(b.speedY)
	if b.rect.bottom() < world.rect.top() {
		b.visible = false
	}

	for i := range world.meteors.items {
		if world.meteors.items[i].collideCircle(b.rect.centerX(), b.rect.centerY(), b.radius) {
			b.visible = false
			world.meteors.items[i].reset(world)
		}
	}
}
