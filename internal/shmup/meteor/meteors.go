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
		m.items[i] = newMeteor(o)
	}
	return m
}

func (m *Meteors) Update() {
	for i := range m.items {
		m.items[i].update()
	}
}

func (m *Meteors) Draw(screen *ebiten.Image) {
	for i := range m.items {
		m.items[i].draw(screen)
	}
}

func (m *Meteors) CollideCircle(c *shape.Circle) bool {
	for i := range m.items {
		if m.items[i].collideCircle(c) {
			m.items[i].reset()
			return true
		}
	}
	return false
}
