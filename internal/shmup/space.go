package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Space struct {
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
}

func newSpace(g *Game) *Space {
	s := &Space{}
	s.image = g.res.images["space"]
	s.options = &ebiten.DrawImageOptions{}
	return s
}

func (s *Space) draw(screen *ebiten.Image) {
	screen.DrawImage(s.image, s.options)
}
