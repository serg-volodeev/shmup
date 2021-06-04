package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Meteors struct {
	items []*Meteor
}

func newMeteors(g *Game) *Meteors {
	count := 8
	m := &Meteors{}
	m.items = make([]*Meteor, count)
	for i := 0; i < count; i++ {
		m.items[i] = newMeteor(g)
	}
	return m
}

func (m *Meteors) update() {
	for i := range m.items {
		m.items[i].update()
	}
}

func (m *Meteors) draw(screen *ebiten.Image) {
	for i := range m.items {
		m.items[i].draw(screen)
	}
}

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
