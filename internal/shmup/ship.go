package shmup

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Ship struct {
	rect    *Rect
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
	speedX  float64
	radius  float64
}

func newShip(res *Res, world *World) *Ship {
	s := &Ship{}
	s.image = res.images["ship"]
	s.options = &ebiten.DrawImageOptions{}
	s.speedX = 4

	s.rect = newRectFromImage(s.image)
	s.rect.setCenterX(world.rect.centerX())
	s.rect.setBottom(world.rect.bottom() - 10)

	s.radius = s.rect.h/2 - 2
	return s
}

func (s *Ship) draw(screen *ebiten.Image) {
	s.options.GeoM.Reset()
	s.options.GeoM.Translate(s.rect.x, s.rect.y)
	screen.DrawImage(s.image, s.options)

	// x := s.rect.centerX() - s.radius
	// y := s.rect.centerY() - s.radius
	// ebitenutil.DrawRect(screen, x, y, s.radius*2, s.radius*2, color.White)
}

func (s *Ship) update(world *World) error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		s.rect.moveX(s.speedX)
		if s.rect.right() > world.rect.right() {
			s.rect.setRight(world.rect.right())
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		s.rect.moveX(-s.speedX)
		if s.rect.left() < world.rect.left() {
			s.rect.setLeft(world.rect.left())
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.fire(world)
	}
	if world.meteors.collideCircle(s.rect.centerX(), s.rect.centerY(), s.radius) {
		return fmt.Errorf("collide meteor")
	}
	return nil
}

func (s *Ship) fire(world *World) {
	b := world.bullets.newBullet()
	b.rect.setCenterX(s.rect.centerX())
	b.rect.setBottom(s.rect.top())
}
