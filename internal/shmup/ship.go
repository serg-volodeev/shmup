package shmup

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/serg-volodeev/shmup/internal/shape"
)

type Ship struct {
	rect    *shape.Rect
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

	s.rect = shape.NewRectFromImage(s.image)
	s.rect.SetCenterX(world.rect.CenterX())
	s.rect.SetBottom(world.rect.Bottom() - 10)

	s.radius = s.rect.Height()/2 - 2
	return s
}

func (s *Ship) draw(screen *ebiten.Image) {
	s.options.GeoM.Reset()
	s.options.GeoM.Translate(s.rect.Left(), s.rect.Top())
	screen.DrawImage(s.image, s.options)

	// x := s.rect.centerX() - s.radius
	// y := s.rect.centerY() - s.radius
	// ebitenutil.DrawRect(screen, x, y, s.radius*2, s.radius*2, color.White)
}

func (s *Ship) update(world *World) error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		s.rect.MoveX(s.speedX)
		if s.rect.Right() > world.rect.Right() {
			s.rect.SetRight(world.rect.Right())
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		s.rect.MoveX(-s.speedX)
		if s.rect.Left() < world.rect.Left() {
			s.rect.SetLeft(world.rect.Left())
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.fire(world)
	}
	if world.meteors.collideCircle(s.rect.CenterX(), s.rect.CenterY(), s.radius) {
		return fmt.Errorf("collide meteor")
	}
	return nil
}

func (s *Ship) fire(world *World) {
	b := world.bullets.newBullet()
	b.rect.SetCenterX(s.rect.CenterX())
	b.rect.SetBottom(s.rect.Top())
}
