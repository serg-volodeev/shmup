package shmup

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/serg-volodeev/shmup/internal/shape"
)

type Meteor struct {
	rect     *shape.Rect
	image    *ebiten.Image
	options  *ebiten.DrawImageOptions
	speedX   float64
	speedY   float64
	rotAngle float64
	rotSpeed float64
	radius   float64
}

func randRange(min, max int32) float64 {
	return float64(rand.Int31n(max-min) + min)
}

func newMeteor(res *Res, world *World) *Meteor {
	imgName := fmt.Sprintf("meteor%d", int(randRange(1, 4)))
	m := &Meteor{}
	m.image = res.images[imgName]
	m.options = &ebiten.DrawImageOptions{}
	m.rect = shape.NewRectFromImage(m.image)
	m.radius = m.rect.Height()/2 - 2
	m.reset(world)
	return m
}

func (m *Meteor) reset(world *World) {
	m.rect.SetLeft(randRange(0, int32(world.rect.Right())-int32(m.rect.Width())))
	m.rect.SetTop(randRange(-100, -40))
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

	// x := m.rect.centerX() - m.radius
	// y := m.rect.centerY() - m.radius
	// ebitenutil.DrawRect(screen, x, y, m.radius*2, m.radius*2, color.White)

}

func (m *Meteor) update(world *World) {
	m.rect.MoveY(m.speedY)
	m.rect.MoveX(m.speedX)

	m.rotAngle += m.rotSpeed
	if m.rotAngle > 360 || m.rotAngle < -360 {
		m.rotAngle = 0
	}

	if m.rect.Top() > world.rect.Bottom()+100 {
		m.reset(world)
	}
}

func square(n float64) float64 {
	return n * n
}

func (m *Meteor) collideCircle(x, y, radius float64) bool {
	mx := m.rect.CenterX()
	my := m.rect.CenterY()
	return square(mx-x)+square(my-y) < square(m.radius+radius)
}
