package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/serg-volodeev/shmup/internal/shape"
)

type Bullet struct {
	game    *Game
	rect    *shape.Rect
	circle  *shape.Circle
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
	speedY  float64
	visible bool
}

func newBullet(g *Game) *Bullet {
	b := &Bullet{}
	b.game = g
	b.image = g.res.GetImage("bullet")
	b.options = &ebiten.DrawImageOptions{}
	b.speedY = -8
	b.visible = true
	b.rect = shape.NewRectFromImage(b.image)
	b.circle = shape.NewCircle(b.rect.CenterX(), b.rect.CenterY(), b.rect.Width()/2)
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

func (b *Bullet) update() {
	if !b.visible {
		return
	}
	b.rect.MoveY(b.speedY)
	b.circle.SetCenter(b.rect.CenterX(), b.rect.CenterY())
	if b.rect.Bottom() < b.game.rect.Top() {
		b.visible = false
	}
	if b.game.meteors.collideBullet(b) {
		b.visible = false
	}
}
