package space

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/serg-volodeev/shmup/internal/shmup/res"
)

type Opts struct {
	Res *res.Res
}

type Space struct {
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
}

func NewSpace(o *Opts) *Space {
	s := &Space{}
	s.image = o.Res.GetImage("space")
	s.options = &ebiten.DrawImageOptions{}
	return s
}

func (s *Space) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.image, s.options)
}
