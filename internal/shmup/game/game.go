package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game interface {
	GetImage(name string) *ebiten.Image
}
