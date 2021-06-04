package shmup

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/serg-volodeev/shmup/internal/shape"
)

type Ship struct {
	game    *Game
	rect    *shape.Rect
	circle  *shape.Circle
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
	speedX  float64
}

func newShip(g *Game) *Ship {
	s := &Ship{}
	s.game = g
	s.image = g.res.images["ship"]
	s.options = &ebiten.DrawImageOptions{}
	s.speedX = 4

	s.rect = shape.NewRectFromImage(s.image)
	s.rect.SetCenterX(g.rect.CenterX())
	s.rect.SetBottom(g.rect.Bottom() - 10)

	s.circle = shape.NewCircle(s.rect.CenterX(), s.rect.CenterY(), s.rect.Height()/2-2)
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

func (s *Ship) update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		s.move(s.speedX)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		s.move(-s.speedX)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.fire()
	}
	if s.game.meteors.collideShip(s) {
		return fmt.Errorf("collide meteor")
	}
	return nil
}

func (s *Ship) move(dx float64) {
	s.rect.MoveX(dx)
	if s.rect.Right() > s.game.rect.Right() {
		s.rect.SetRight(s.game.rect.Right())
	}
	if s.rect.Left() < s.game.rect.Left() {
		s.rect.SetLeft(s.game.rect.Left())
	}
	s.circle.SetCenter(s.rect.CenterX(), s.rect.CenterY())
}

func (s *Ship) fire() {
	b := s.game.bullets.newBullet()
	b.rect.SetCenterX(s.rect.CenterX())
	b.rect.SetBottom(s.rect.Top())
}
