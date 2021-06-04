package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/serg-volodeev/shmup/internal/shape"
)

type Bullet struct {
	rect    *shape.Rect
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
	speedY  float64
	visible bool
	radius  float64
}

func newBullet(g *Game) *Bullet {
	b := &Bullet{}
	b.image = g.res.images["bullet"]
	b.options = &ebiten.DrawImageOptions{}
	b.speedY = -8
	b.visible = true
	b.rect = shape.NewRectFromImage(b.image)
	b.radius = b.rect.Width() / 2
	return b
}

func (b *Bullet) draw(screen *ebiten.Image) {
	if !b.visible {
		return
	}
	b.options.GeoM.Reset()
	b.options.GeoM.Translate(b.rect.Left(), b.rect.Top())
	screen.DrawImage(b.image, b.options)
}

func (b *Bullet) update(g *Game) {
	if !b.visible {
		return
	}
	b.rect.MoveY(b.speedY)
	if b.rect.Bottom() < g.rect.Top() {
		b.visible = false
	}

	for i := range g.meteors.items {
		if g.meteors.items[i].collideCircle(b.rect.CenterX(), b.rect.CenterY(), b.radius) {
			b.visible = false
			g.meteors.items[i].reset(g)
		}
	}
}
