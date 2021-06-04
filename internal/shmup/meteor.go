package shmup

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/serg-volodeev/shmup/internal/shape"
)

type Meteor struct {
	game     *Game
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

func newMeteor(g *Game) *Meteor {
	imgName := fmt.Sprintf("meteor%d", int(randRange(1, 4)))
	m := &Meteor{}
	m.game = g
	m.image = g.res.images[imgName]
	m.options = &ebiten.DrawImageOptions{}
	m.rect = shape.NewRectFromImage(m.image)
	m.circle = shape.NewCircle(m.rect.CenterX(), m.rect.CenterY(), m.rect.Height()/2-2)
	m.reset()
	return m
}

func (m *Meteor) reset() {
	m.rect.SetLeft(randRange(0, int32(m.game.rect.Right())-int32(m.rect.Width())))
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

	// x := m.rect.centerX() - m.radius
	// y := m.rect.centerY() - m.radius
	// ebitenutil.DrawRect(screen, x, y, m.radius*2, m.radius*2, color.White)

}

func (m *Meteor) update() {
	m.rect.Move(m.speedX, m.speedY)
	m.circle.SetCenter(m.rect.CenterX(), m.rect.CenterY())

	m.rotAngle += m.rotSpeed
	if m.rotAngle > 360 || m.rotAngle < -360 {
		m.rotAngle = 0
	}

	if m.rect.Top() > m.game.rect.Bottom()+100 {
		m.reset()
	}
}

func (m *Meteor) collideCircle(circle *shape.Circle) bool {
	return m.circle.CollideCircle(circle)
}
