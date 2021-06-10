package meteor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/serg-volodeev/shmup/internal/shape"
)

type Meteors struct {
	items []*Meteor
}

func NewMeteors(o *Opts) *Meteors {
	count := 8
	m := &Meteors{}
	m.items = make([]*Meteor, count)
	for i := 0; i < count; i++ {
		m.items[i] = NewMeteor(o)
	}
	return m
}

func (m *Meteors) Update() {
	for i := range m.items {
		m.items[i].Update()
	}
}

func (m *Meteors) Draw(screen *ebiten.Image) {
	for i := range m.items {
		m.items[i].Draw(screen)
	}
}

func (m *Meteors) CollideCircle(c *shape.Circle) bool {
	for i := range m.items {
		if m.items[i].CollideCircle(c) {
			m.items[i].Reset()
			return true
		}
	}
	return false
}

/*
func (m *Meteors) collideBullet(b *Bullet) bool {
	for i := range m.items {
		if m.items[i].collideCircle(b.circle) {
			m.items[i].reset()
			return true
		}
	}
	return false
}

func (m *Meteors) collideShip(s *Ship) bool {
	for i := range m.items {
		if m.items[i].collideCircle(s.circle) {
			return true
		}
	}
	return false
}
*/
