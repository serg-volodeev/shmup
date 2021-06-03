package shmup

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Meteor struct {
	rect     *Rect
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
	m.rect = newRectFromImage(m.image)
	m.radius = m.rect.h/2 - 2
	m.reset(world)
	return m
}

func (m *Meteor) reset(world *World) {
	m.rect.x = randRange(0, int32(world.rect.right())-int32(m.rect.w))
	m.rect.y = randRange(-100, -40)
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
	m.options.GeoM.Translate(m.rect.centerX(), m.rect.centerY())
	screen.DrawImage(m.image, m.options)

	// x := m.rect.centerX() - m.radius
	// y := m.rect.centerY() - m.radius
	// ebitenutil.DrawRect(screen, x, y, m.radius*2, m.radius*2, color.White)

}

func (m *Meteor) update(world *World) {
	m.rect.moveY(m.speedY)
	m.rect.moveX(m.speedX)

	m.rotAngle += m.rotSpeed
	if m.rotAngle > 360 || m.rotAngle < -360 {
		m.rotAngle = 0
	}

	if m.rect.top() > world.rect.bottom()+100 {
		m.reset(world)
	}
}

func square(n float64) float64 {
	return n * n
}

func (m *Meteor) collideCircle(x, y, radius float64) bool {
	mx := m.rect.centerX()
	my := m.rect.centerY()
	return square(mx-x)+square(my-y) < square(m.radius+radius)
}
