package bullet

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/serg-volodeev/shmup/internal/shape"
	"github.com/serg-volodeev/shmup/internal/shmup/meteor"
	"github.com/serg-volodeev/shmup/internal/shmup/res"
)

type Opts struct {
	Bounds  *shape.Rect
	Res     *res.Res
	Meteors *meteor.Meteors
}

type Bullet struct {
	bounds  *shape.Rect
	meteors *meteor.Meteors
	rect    *shape.Rect
	circle  *shape.Circle
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
	speedY  float64
	visible bool
}

func newBullet(o *Opts) *Bullet {
	b := &Bullet{}
	b.bounds = o.Bounds
	b.meteors = o.Meteors
	b.image = o.Res.GetImage("bullet")
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
	if b.rect.Bottom() < b.bounds.Top() {
		b.visible = false
	}
	if b.meteors.CollideCircle(b.circle) {
		b.visible = false
	}
}
