package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Ship struct {
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
	x, y    float64
	speedX  float64
}

func newShip(res *Res) *Ship {
	sh := &Ship{}
	sh.image = res.images["ship"]
	sh.x = float64(ScreenWidth)/2 - sh.width()/2
	sh.y = float64(ScreenHeight) - sh.height() - 10
	sh.options = &ebiten.DrawImageOptions{}
	sh.speedX = 4
	return sh
}

func (sh *Ship) width() float64 {
	return float64(sh.image.Bounds().Dx())
}

func (sh *Ship) height() float64 {
	return float64(sh.image.Bounds().Dy())
}

func (sh *Ship) draw(screen *ebiten.Image) {
	sh.options.GeoM.Reset()
	sh.options.GeoM.Translate(sh.x, sh.y)
	screen.DrawImage(sh.image, sh.options)
}

func (sh *Ship) update() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && (sh.x+sh.width()) < ScreenWidth {
		sh.x += sh.speedX
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && sh.x > 0 {
		sh.x -= sh.speedX
	}
}
