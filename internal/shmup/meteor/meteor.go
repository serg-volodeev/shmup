package meteor

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/serg-volodeev/shmup/internal/shape"
	"github.com/serg-volodeev/shmup/internal/shmup/res"
)

type Opts struct {
	Bounds *shape.Rect
	Res    *res.Res
}

type Meteor struct {
	bounds   *shape.Rect
	rect     *shape.Rect
	circle   *shape.Circle
	image    *ebiten.Image
	options  *ebiten.DrawImageOptions
	speedX   float64
	speedY   float64
	rotAngle float64
	rotSpeed float64
}

func randRange(min, max int32) float64 {
	return float64(rand.Int31n(max-min) + min)
}

func newMeteor(o *Opts) *Meteor {
	imgName := fmt.Sprintf("meteor%d", int(randRange(1, 4)))
	m := &Meteor{}
	m.bounds = o.Bounds
	m.image = o.Res.GetImage(imgName)
	m.options = &ebiten.DrawImageOptions{}
	m.rect = shape.NewRectFromImage(m.image)
	m.circle = shape.NewCircle(m.rect.CenterX(), m.rect.CenterY(), m.rect.Height()/2-2)
	m.reset()
	return m
}

func (m *Meteor) reset() {
	m.rect.SetLeft(randRange(0, int32(m.bounds.Right())-int32(m.rect.Width())))
	m.rect.SetTop(randRange(-100, -40))
	m.circle.SetCenter(m.rect.CenterX(), m.rect.CenterY())
	m.speedY = randRange(1, 8)
	m.speedX = randRange(-3, 3)
	m.rotAngle = 0
	m.rotSpeed = randRange(-8, 8) / 100
}

func (m *Meteor) draw(screen *ebiten.Image) {
	w, h := m.image.Size()
	m.options.GeoM.Reset()
	m.options.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	m.options.GeoM.Rotate(m.rotAngle)
	m.options.GeoM.Translate(m.rect.CenterX(), m.rect.CenterY())
	screen.DrawImage(m.image, m.options)
}

func (m *Meteor) update() {
	m.rect.Move(m.speedX, m.speedY)
	m.circle.SetCenter(m.rect.CenterX(), m.rect.CenterY())

	m.rotAngle += m.rotSpeed
	if m.rotAngle > 360 || m.rotAngle < -360 {
		m.rotAngle = 0
	}

	if m.rect.Top() > m.bounds.Bottom()+100 {
		m.reset()
	}
}

func (m *Meteor) collideCircle(c *shape.Circle) bool {
	return m.circle.CollideCircle(c)
}
