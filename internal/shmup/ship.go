package shmup

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ship struct {
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
}

func NewShip() *Ship {
	ship := &Ship{}
	ship.image = ebiten.NewImage(32, 32)
	ship.image.Fill(color.Black)
	ship.options = &ebiten.DrawImageOptions{}
	ship.options.GeoM.Translate(10, 20)
	return ship
}

func (ship *Ship) Draw(screen *ebiten.Image) {
	screen.DrawImage(ship.image, ship.options)
}
