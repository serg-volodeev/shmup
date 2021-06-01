package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Space struct {
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
}

func NewSpace() *Space {
	space := &Space{}
	space.image, _ = LoadImageFromFile("./assets/images/starfield.png")
	space.options = &ebiten.DrawImageOptions{}
	return space
}

func (space *Space) Draw(screen *ebiten.Image) {
	screen.DrawImage(space.image, space.options)
}
