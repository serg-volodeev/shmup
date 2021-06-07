package space

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/serg-volodeev/shmup/internal/shmup/game"
)

type Space struct {
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
}

func NewSpace(g game.Game) *Space {
	s := &Space{}
	s.image = g.GetImage("space")
	s.options = &ebiten.DrawImageOptions{}
	return s
}

func (s *Space) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.image, s.options)
}
