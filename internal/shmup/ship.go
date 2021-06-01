package shmup

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Ship struct {
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
	x, y    float64
	speedX  float64
}

func NewShip() *Ship {
	ship := &Ship{}
	ship.image, _ = LoadImageFromFile("./assets/images/playerShip1_orange.png")
	ship.x = float64(ScreenWidth)/2 - ship.width()/2
	ship.y = float64(ScreenHeight) - ship.height() - 10
	ship.options = &ebiten.DrawImageOptions{}
	ship.speedX = 4
	return ship
}

func (ship *Ship) width() float64 {
	return float64(ship.image.Bounds().Dx())
}

func (ship *Ship) height() float64 {
	return float64(ship.image.Bounds().Dy())
}

func (ship *Ship) Draw(screen *ebiten.Image) {
	ship.options.GeoM.Reset()
	ship.options.GeoM.Translate(ship.x, ship.y)
	screen.DrawImage(ship.image, ship.options)
}

func (ship *Ship) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && (ship.x+ship.width()) < ScreenWidth {
		ship.x += ship.speedX
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && ship.x > 0 {
		ship.x -= ship.speedX
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		ship.fire()
	}
}

func (ship *Ship) fire() {
	b := NewBullet()
	b.x = ship.x + ship.width()/2 - b.width()/2
	b.y = ScreenHeight - ship.height() - 10 - b.height()
}
