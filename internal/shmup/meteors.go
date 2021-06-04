package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/serg-volodeev/shmup/internal/shape"
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

func (m *Meteors) update(g *Game) {
	for i := range m.items {
		m.items[i].update(g)
	}
}

func (m *Meteors) draw(screen *ebiten.Image) {
	for i := range m.items {
		m.items[i].draw(screen)
	}
}

func (m *Meteors) collideCircle(circle *shape.Circle) bool {
	for i := range m.items {
		if m.items[i].collideCircle(circle) {
			return true
		}
	}
	return false
}
