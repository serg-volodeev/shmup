package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
	x, y    float64
	speedY  float64
	visible bool
}

func newBullet(res *Res) *Bullet {
	b := &Bullet{}
	b.image = res.images["bullet"]
	b.options = &ebiten.DrawImageOptions{}
	b.speedY = -8
	b.visible = true
	return b
}

func (b *Bullet) width() float64 {
	return float64(b.image.Bounds().Dx())
}

func (b *Bullet) height() float64 {
	return float64(b.image.Bounds().Dy())
}

func (b *Bullet) draw(screen *ebiten.Image) {
	if !b.visible {
		return
	}
	b.options.GeoM.Reset()
	b.options.GeoM.Translate(b.x, b.y)
	screen.DrawImage(b.image, b.options)
}

func (b *Bullet) update() {
	if !b.visible {
		return
	}
	b.y += b.speedY
	if b.y+b.height() < 0 {
		b.visible = false
	}
}
