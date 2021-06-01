package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var bulletImage *ebiten.Image
var bullets []*Bullet

type Bullet struct {
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
	x, y    float64
	speedY  float64
	visible bool
}

func NewBullet() *Bullet {
	if bulletImage == nil {
		bulletImage, _ = LoadImageFromFile("./assets/images/laserRed16.png")
	}

	if bullets == nil {
		bullets = make([]*Bullet, 0, 20)
	}

	for i := range bullets {
		if !bullets[i].visible {
			bullets[i].visible = true
			return bullets[i]
		}
	}

	b := &Bullet{}
	b.image = bulletImage
	b.options = &ebiten.DrawImageOptions{}
	b.speedY = -8
	b.visible = true
	bullets = append(bullets, b)
	return b
}

func (b *Bullet) width() float64 {
	return float64(b.image.Bounds().Dx())
}

func (b *Bullet) height() float64 {
	return float64(b.image.Bounds().Dy())
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	b.options.GeoM.Reset()
	b.options.GeoM.Translate(b.x, b.y)
	screen.DrawImage(b.image, b.options)
}

func (b *Bullet) Update() {
	if !b.visible {
		return
	}
	b.y += b.speedY
	if b.y+b.height() < 0 {
		b.visible = false
	}
}

func updateBullets() {
	for i := range bullets {
		bullets[i].Update()
	}
}

func drawBullets(screen *ebiten.Image) {
	for i := range bullets {
		bullets[i].Draw(screen)
	}
}
